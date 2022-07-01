package models

import (
	"log"
	"sync"

	"github.com/gomodule/redigo/redis"
)

type RedisDB struct {
	redis.Conn
}

var RedisCN = GetConn()

var (
	once sync.Once
	p    *redis.Pool
)

func GetConn() *redis.Pool {
	p.MaxIdle = 20
	p.MaxActive = 20
	p.IdleTimeout = 240
	once.Do(func() {
		p = &redis.Pool{
			Dial: func() (redis.Conn, error) {
				conn, err := redis.Dial("tcp", "redis-master:6379")
				if err != nil {
					log.Fatal("ERROR: No se puede conectar con Redis")
				}

				if _, err_2 := conn.Do("AUTH", "dfgfgq4356qdfgawet52q345"); err_2 != nil {
					conn.Close()
					return nil, err_2
				}

				return conn, err
			},
		}
	})

	return p
}
