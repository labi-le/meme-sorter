package web

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"log"
	"meme-sorter/internal/structures"
	"net/http"
	"time"
)

type Server struct {
	router *mux.Router
	logger *logrus.Logger

	Config *structures.Config
}

// implement
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func newServer(config *structures.Config) *Server {
	s := &Server{
		router: mux.NewRouter(),
		logger: logrus.New(),

		Config: config,
	}

	s.route()

	return s
}

func Start(config *structures.Config) error {
	srv := newServer(config)
	srv.configureLogger()

	srv.logger.Log(logrus.InfoLevel, "Rest api started")

	server := &http.Server{
		Handler: srv,
		Addr:    config.Addr,
	}

	log.Fatal(server.ListenAndServe())

	return server.ListenAndServe()
}

func (s *Server) configureLogger() {
	level, err := logrus.ParseLevel(s.Config.LogLevel)
	if err != nil {
		panic("invalid log level")
	}

	s.logger.SetLevel(level)
}

func (s *Server) route() {
	s.router.Use(s.logRequestMiddleware)
	s.router.HandleFunc("/api/{method}", s.apiResolver).Methods(http.MethodPost)
}

func (s *Server) logRequestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := s.logger.WithFields(logrus.Fields{
			"IP": r.RemoteAddr,
		})
		logger.Infof("started %s %s", r.Method, r.RequestURI)

		start := time.Now()
		rw := &responseWriter{w, http.StatusOK}
		next.ServeHTTP(rw, r)

		var level logrus.Level
		switch {
		case rw.code >= 500:
			level = logrus.ErrorLevel
		case rw.code >= 400:
			level = logrus.WarnLevel
		default:
			level = logrus.InfoLevel
		}
		logger.Logf(
			level,
			"completed with %d %s in %v",
			rw.code,
			http.StatusText(rw.code),
			time.Now().Sub(start),
		)
	})
}

func (s *Server) apiResolver(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var Item structures.Meme
	body, _ := ioutil.ReadAll(r.Body)

	err := json.Unmarshal(body, &Item)
	if err != nil {
		response(structures.Response{
			Status:      structures.Failed,
			Description: err.Error(),
			Data:        []string{},
		}, w)
	}

	params := mux.Vars(r)

	var MethodResponse structures.Response

	method := NewMethod(s.Config.DB, &Item)
	switch params["method"] {
	case "create":
		MethodResponse = method.Create()
	case "update":
		MethodResponse = method.Update()
	case "take":
		MethodResponse = method.Read()
	case "delete":
		MethodResponse = method.Delete()
	}

	response(MethodResponse, w)

}

func response(response structures.Response, w http.ResponseWriter) {
	switch response.Status {
	case structures.Success:
		w.WriteHeader(http.StatusBadRequest)
		break

	case structures.Partially:
		w.WriteHeader(http.StatusPartialContent)
		break

	case structures.Failed:
		w.WriteHeader(http.StatusOK)
		break
	}

	_ = json.NewEncoder(w).Encode(response)
	return
}