package models

import (
	"log"
	"time"

	"github.com/gomodule/redigo/redis"
)

type RedisDB struct {
	redis.Conn
}

var RedisCN = GetConn()

func GetConn() *redis.Pool {

	p := &redis.Pool{
		MaxIdle:         20,
		MaxActive:       120,
		IdleTimeout:     3 * time.Second,
		MaxConnLifetime: 10 * time.Second,
		Wait:            true,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", "redis:6379")
			if err != nil {
				log.Fatal("ERROR: No se puede conectar con Redis")
			}
			return conn, err
		},
	}

	return p
}
