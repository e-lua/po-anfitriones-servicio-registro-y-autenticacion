package models

import (
	"log"
	"sync"

	"github.com/gomodule/redigo/redis"
)

type RedisDB_Slave struct {
	redis.Conn
}

var RedisCN_Slave = GetConn_Slave()

var (
	once_slave sync.Once
	p_slave    *redis.Pool
)

func GetConn_Slave() *redis.Pool {
	once_slave.Do(func() {
		p_slave = &redis.Pool{
			Dial: func() (redis.Conn, error) {
				conn, err := redis.Dial("tcp", "redis:6380")
				if err != nil {
					log.Fatal("ERROR: No se puede conectar con Redis")
				}
				return conn, err
			},
		}
	})

	return p_slave
}
