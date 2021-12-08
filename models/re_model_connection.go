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
				conn, err := redis.Dial("tcp", "rediss://default:3eBnVvBJMUpeJdaq@po-comensales-anfitriones-do-user-10365906-0.b.db.ondigitalocean.com:25061")
				if err != nil {
					log.Fatal("ERROR: No se puede conectar con Redis")
				}
				return conn, err
			},
		}
	})

	return p
}
