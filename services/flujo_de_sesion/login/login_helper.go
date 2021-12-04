package login

type Response struct {
	Error     bool   `json:"error"`
	DataError string `json:"dataError"`
	Data      string `json:"data"`
}

type ResponseJWT struct {
	Error     bool       `json:"error"`
	DataError string     `json:"dataError"`
	Data      JWTRequest `json:"data"`
}

type JWT struct {
	IdBusiness int `json:"idBusiness"`
	IdWorker   int `json:"idWorker"`
	IdCountry  int `json:"country"`
	IdRol      int `json:"rol"`
}

type JWTRequest struct {
	IdBusiness int    `json:"idBusiness"`
	IdWorker   int    `json:"idWorker"`
	IdCountry  int    `json:"country"`
	IdRol      int    `json:"rol"`
	Service    string `json:"service"`
	Module     string `json:"module"`
	Epic       string `json:"epic"`
	Endpoint   string `json:"endpoint"`
}
