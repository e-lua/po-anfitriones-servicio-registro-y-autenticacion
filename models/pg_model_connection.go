package models

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

var PostgresCN = Conectar_Pg_DB()

var (
	once_pg sync.Once
	p_pg    *pgxpool.Pool
)

func Conectar_Pg_DB() *pgxpool.Pool {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	p.MaxIdle = 10
	p.MaxActive = 10
	p.IdleTimeout = 240
	once_pg.Do(func() {
		urlString := "postgres://postgresx:adsfg465WFVFGdsrf3465QWFDSFGH4fsadf4fwedf@postgres:5432/postgresx?pool_max_conns=200"

		config, error_connec_pg := pgxpool.ParseConfig(urlString)

		if error_connec_pg != nil {
			log.Fatal("Error en el servidor interno en el driver de PostgreSQL, mayor detalle: " + error_connec_pg.Error())
		}

		p_pg, _ = pgxpool.ConnectConfig(ctx, config)
	})

	return p_pg
}
