package registro

type Response_WithInt struct {
	Error     bool   `json:"error"`
	DataError string `json:"dataError"`
	Data      int    `json:"data"`
}

type Response_SignInFirstStep struct {
	Error     bool            `json:"error"`
	DataError string          `json:"dataError"`
	Data      SignInFirstStep `json:"data"`
}

type SignInFirstStep struct {
	Phone   int `json:"phone"`
	Country int `json:"country"`
}

type Response_WithString struct {
	Error     bool   `json:"error"`
	DataError string `json:"dataError"`
	Data      string `json:"data"`
}

type Response_WithPhoneCountryCode struct {
	Error     bool             `json:"error"`
	DataError string           `json:"dataError"`
	Data      PhoneCountryCode `json:"data"`
}

type PhoneCountryCode struct {
	Phone   int `json:"phone"`
	Country int `json:"country"`
	Code    int `json:"code"`
}

type EntryData_Password struct {
	NewPassword string `json:"newpassword"`
	Phone       int    `json:"phone"`
	Country     int    `json:"country"`
	Code        int    `json:"code"`
}
