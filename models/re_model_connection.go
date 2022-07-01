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
	once.Do(func() {
		p = &redis.Pool{
			Dial: func() (redis.Conn, error) {
				conn, err := redis.Dial("tcp", "redis-master:6379")
				if err != nil {
					log.Fatal("ERROR: No se puede conectar con Redis")
				}
				/*conn, err := redis.DialURL("redis://dfgfgq4356qdfgawet52q345@redis-master:6379")
				if err != nil {
					log.Fatal("ERROR: No se puede conectar con Redis")
				}*/

				return conn, err
			},
		}
	})

	return p
}
