package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host      = "postgres"
	port      = "5432"
	user      = "postgresx"
	password  = "postgresx"
	dbname_pg = "postgresx"
)

var PostgresCN = Conectar_Pg_DB()

func Conectar_Pg_DB() *sql.DB {

	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, user, dbname_pg, password, port)
	db, err := sql.Open("postgres", dbURI)
	if err != nil {
		log.Fatal("Error en el servidor interno en el driver de PostgreSQL, mayor detalle: " + err.Error())
	}
	if err = db.Ping(); err != nil {
		log.Fatal("Error en el servidor interno al intentar conectarse con la base de datos, mayor detalle: " + err.Error())
	}

	//Conexión corecta
	log.Printf("Conexión exitosa con la BD pg_")

	return db
}

//ChequeoConnection es el Ping a la BD
func ChequeoConnection_Pg() int {

	err := PostgresCN.Ping()
	if err != nil {
		return 0
	}
	return 1

}
