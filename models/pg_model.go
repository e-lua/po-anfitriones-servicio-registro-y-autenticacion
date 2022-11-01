package models

import "time"

type Pg_BusinessWorker struct {
	IdBusiness   int       `json:"idBusiness"`
	IdWorker     int       `json:"idWorker"`
	IdCountry    int       `json:"country"`
	CodeRedis    int       `json:"code"`
	Name         string    `json:"name"`
	IdRol        int       `json:"rol"`
	LastName     string    `json:"lastName"`
	Phone        int       `json:"phone"`
	Password     string    `json:"password"`
	Isbanned     bool      `json:"isbanned"`
	SessionCode  int       `json:"sessioncode"`
	UpdatedDate  time.Time `json:"updateDate"`
	IsDeleted    bool      `json:"isdeleted"`
	Email        string    `json:"email"`
	IDDevice     string    `json:"iddevice"`
	IsSubsidiary bool      `json:"issubsidiary"`
	SubsidiaryOf int       `json:"subsidiaryof"`
}

type Pg_SubWorker struct {
	IdWorker       int    `json:"idWorker"`
	IdBusiness     int    `json:"idBusiness"`
	IdRol          int    `json:"rol"`
	Name           string `json:"name"`
	LastName       string `json:"lastName"`
	IdCountry      int    `json:"country"`
	Phone          int    `json:"phone"`
	DateRegistered string `json:"dateregistered"`
}

type V2_Pg_SubWorker struct {
	IdWorker   int    `json:"idWorker"`
	IdBusiness int    `json:"idBusiness"`
	IdRol      int    `json:"rol"`
	Name       string `json:"name"`
	LastName   string `json:"lastName"`
	Email      string `json:"email"`
}
