package models

import (
	"github.com/gomodule/redigo/redis"
)

type RedisDB struct {
	redis.Conn
}

var RedisCN = GetConn()

func GetConn() redis.Conn {

	localURL := "redis:6379"
	//redisURL := "rediss://" + username_re + ":" + password_re + "@" + host_re + ":" + port_re
	redisURL := "redis://" + localURL
	c, err := connectRedis(redisURL)

	if err != nil {
		return nil
	}

	return c
}

//format url ==> "redis://redis:6379"
func connectRedis(redisURL string) (redis.Conn, error) {
	if redisURL != "" {
		redisPassword := ""
		return redis.DialURL(redisURL, redis.DialPassword(redisPassword))
	} else {
		localURL := "redis:6379"
		//redisURL := "rediss://" + username_re + ":" + password_re + "@" + host_re + ":" + port_re
		redisURL := "redis://" + localURL
		return redis.DialURL(redisURL)
	}
}
