package models

import "time"

type Pg_BusinessWorker struct {
	IdBusiness  int       `json:"idBusiness"`
	IdWorker    int       `json:"idWorker"`
	IdCountry   int       `json:"country"`
	CodeRedis   int       `json:"code"`
	Name        string    `json:"name"`
	IdRol       int       `json:"rol"`
	LastName    string    `json:"lastName"`
	Phone       int       `json:"phone"`
	Password    string    `json:"password"`
	Isbanned    bool      `json:"isbanned"`
	SessionCode int       `json:"sessioncode"`
	UpdatedDate time.Time `json:"updateDate"`
}
