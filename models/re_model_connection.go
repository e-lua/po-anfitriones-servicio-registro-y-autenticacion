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
		MaxIdle:         50,
		MaxActive:       120,
		IdleTimeout:     4 * time.Second,
		MaxConnLifetime: 4 * time.Second,
		Wait:            true,
		Dial: func() (redis.Conn, error) {
			redis.DialConnectTimeout(2 * time.Second)
			redis.DialReadTimeout(2 * time.Second)
			redis.DialWriteTimeout(2 * time.Second)
			conn, err := redis.Dial("tcp", "redis:6379")
			if err != nil {
				log.Fatal("ERROR: No se puede conectar con Redis")
			}

			return conn, err
		},
	}

	return p
}
