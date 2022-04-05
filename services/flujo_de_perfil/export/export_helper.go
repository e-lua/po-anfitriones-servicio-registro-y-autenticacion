package export

type Response_ListString struct {
	Error     bool     `json:"error"`
	DataError string   `json:"dataError"`
	Data      []string `json:"data"`
}
type Request_Export_Notifications struct {
	Idbusiness      int   `json:"idbusiness"`
	Type            int   `json:"type"`
	ArrayBusinesses []int `json:"manybusinesses"`
}
