package models

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

var PostgresCN = Conectar_Pg_DB()

func Conectar_Pg_DB() *pgxpool.Pool {

	urlString := "postgres://postgresx:postgresx@postgres:5432/postgresx?pool_max_conns=50"

	config, error_connec_pg := pgxpool.ParseConfig(urlString)

	if error_connec_pg != nil {
		log.Fatal("Error en el servidor interno en el driver de PostgreSQL, mayor detalle: " + error_connec_pg.Error())
		return nil
	}

	conn, _ := pgxpool.ConnectConfig(context.Background(), config)

	return conn
}
