package models

type Re_SetGetCode struct {
	PhoneRegister_Key int `json:"phoneRegister"`
	Code              int `json:"code"`
	Country           int `json:"country"`
}
