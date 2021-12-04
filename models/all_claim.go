package models

import (
	//jwt se comporta como Alias de lo que esta entre comillas
	jwt "github.com/dgrijalva/jwt-go"
)

//Claim es la estructura usada para procesar el JWT
type Claim struct {
	Phone    int `json:"phone"`
	Business int `json:"business"`
	Worker   int `json:"worker"`
	Country  int `json:"country"`
	IDRol    int ` json:"rol"`
	// uno de los StandardClaims -> fecha de expiración de un token
	jwt.StandardClaims
}
