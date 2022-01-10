package profile

type Response_WithInt struct {
	Error     bool   `json:"error"`
	DataError string `json:"dataError"`
	Data      int    `json:"data"`
}

type Response_WithString struct {
	Error     bool   `json:"error"`
	DataError string `json:"dataError"`
	Data      string `json:"data"`
}

type ResponseJWT struct {
	Error     bool   `json:"error"`
	DataError string `json:"dataError"`
	Data      JWT    `json:"data"`
}

type JWT struct {
	IdBusiness int `json:"idBusiness"`
	IdWorker   int `json:"idWorker"`
	IdCountry  int `json:"country"`
	IdRol      int `json:"rol"`
}

type EntryData_Password struct {
	OldPassword string `json:"oldpassword"`
	NewPassword string `json:"newpassword"`
	Phone       int    `json:"phone"`
	Country     int    `json:"country"`
}

type Entry_Profile struct {
	Name     string `json:"name"`
	LastName string `json:"lastname"`
}
