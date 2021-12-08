package models

import (
	"log"

	"github.com/gomodule/redigo/redis"
)

type RedisDB struct {
	redis.Conn
}

var RedisCN = GetConn()

func GetConn() *redis.Pool {

	pool := &redis.Pool{
		MaxIdle:         20,
		MaxActive:       80,
		IdleTimeout:     2,
		MaxConnLifetime: 10,
		Wait:            true,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", "redis:6379")
			if err != nil {
				log.Fatal("ERROR: No se puede conectar con Redis")
			}
			return conn, err
		},
	}

	return pool
}
