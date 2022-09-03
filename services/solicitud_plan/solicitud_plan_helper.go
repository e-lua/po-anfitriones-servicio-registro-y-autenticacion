package solicitud_plan

type ResponseJWT_Anfitrion struct {
	Error     bool          `json:"error"`
	DataError string        `json:"dataError"`
	Data      JWT_Anfitrion `json:"data"`
}

type JWT_Anfitrion struct {
	IdBusiness int `json:"idBusiness"`
	IdWorker   int `json:"idWorker"`
	IdCountry  int `json:"country"`
	IdRol      int `json:"rol"`
}

type Response struct {
	Error     bool   `json:"error"`
	DataError string `json:"dataError"`
	Data      string `json:"data"`
}
