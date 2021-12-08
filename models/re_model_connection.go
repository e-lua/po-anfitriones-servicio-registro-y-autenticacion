package models

import (
	"log"
	"sync"
	"time"

	"github.com/gomodule/redigo/redis"
)

type RedisDB struct {
	redis.Conn
}

var RedisCN = GetConn()

func GetConn() *redis.Pool {

	max := 120

	p := &redis.Pool{
		MaxIdle:         20,
		MaxActive:       max,
		IdleTimeout:     3 * time.Millisecond,
		MaxConnLifetime: 10 * time.Millisecond,
		Wait:            true,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", "redis:6379")
			if err != nil {
				log.Fatal("ERROR: No se puede conectar con Redis")
			}
			return conn, err
		},
	}

	var wg sync.WaitGroup
	wg.Add(2 * max)
	done := make(chan struct{})
	for i := 0; i < max; i++ {
		c1 := p.Get()
		log.Println("active conn 1:", p.ActiveCount())
		go work(done, c1, &wg)
		c2 := p.Get()
		log.Println("active conn 2:", p.ActiveCount())
		go work(done, c2, &wg)
	}

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		log.Println("requesting connection...")
		now := time.Now()
		for i := 0; i < max; i++ {
			c1 := p.Get()
			log.Println("got connection 1 took:", time.Since(now), "active:", p.ActiveCount())
			c1.Close()
			c2 := p.Get()
			log.Println("got connection 2  took:", time.Since(now), "active:", p.ActiveCount())
			c2.Close()
		}
	}(&wg)
	time.Sleep(time.Second * 2)
	close(done)

	wg.Wait()
	log.Println("active:", p.ActiveCount())

	return p
}

func work(done <-chan struct{}, c redis.Conn, wg *sync.WaitGroup) {
	log.Println("work...")
	defer func() {
		c.Close()
		log.Println("work done")
		wg.Done()
	}()
	<-done
}
