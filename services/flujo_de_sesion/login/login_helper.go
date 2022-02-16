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

/*-------------------------------------*/

type Responde_JWTAndRol struct {
	Error     bool      `json:"error"`
	DataError string    `json:"dataError"`
	Data      JWTAndRol `json:"data"`
}

type JWTAndRol struct {
	JWT      string `json:"jwt"`
	Rol      int    `json:"rol"`
	Phone    int    `json:"phone"`
	Country  int    `json:"country"`
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	ID       int    `json:"id"`
}

/*-------------------------------------*/

type Input_BusinessWorker_login struct {
	Phone           int    `json:"phone"`
	IdCountry       int    `json:"country"`
	Password        string `json:"password"`
	Email           string `json:"email"`
	IsAnfitrion     bool   `json:"isworker"`
	DateTimeSession string `json:"datetime"`
}
