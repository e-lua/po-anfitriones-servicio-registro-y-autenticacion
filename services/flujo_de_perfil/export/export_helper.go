package export

type Response_ListString struct {
	Error     bool     `json:"error"`
	DataError string   `json:"dataError"`
	Data      []string `json:"data"`
}
