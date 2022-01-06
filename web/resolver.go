package web

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"meme-sorter/internal"
	"meme-sorter/internal/structures"
	"net/http"
)

type Resolver struct {
	writer http.ResponseWriter
	reader *http.Request
}

func NewResolver(db *internal.DB, w http.ResponseWriter, r *http.Request) {
	var Item structures.Meme

	w.Header().Set("Content-Type", "application/json")

	body, _ := ioutil.ReadAll(r.Body)

	resolver := Resolver{w, r}

	resolver.errHandler(json.Unmarshal(body, &Item))
	resolver.resolve(db, Item)
}

func (r *Resolver) errHandler(err error) {
	if err != nil {
		r.response(structures.Response{
			Status:      structures.Failed,
			Description: err.Error(),
			Data:        []string{},
		})
	}
}

func (r *Resolver) resolve(d *internal.DB, item structures.Meme) {
	params := mux.Vars(r.reader)

	var MethodResponse structures.Response

	method := NewMethod(d, &item)
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

	r.response(MethodResponse)

}

func (r *Resolver) response(response structures.Response) {
	switch response.Status {
	case structures.Success:
		r.writer.WriteHeader(http.StatusBadRequest)
		break

	case structures.Partially:
		r.writer.WriteHeader(http.StatusPartialContent)
		break

	case structures.Failed:
		r.writer.WriteHeader(http.StatusOK)
		break
	}

	_ = json.NewEncoder(r.writer).Encode(response)
	return
}
