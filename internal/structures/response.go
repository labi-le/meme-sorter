package structures

type Response struct {
	Status      string      `json:"status"`
	Description string      `json:"description,omitempty"`
	Data        interface{} `json:"data"`
}
