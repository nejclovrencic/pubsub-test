package db

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

var pool *redis.Pool

func Init(addr string) {
	pool = &redis.Pool{
		MaxIdle:     100,
		MaxActive:   300,
		IdleTimeout: 240 * time.Second,
		// Dial or DialContext must be set. When both are set, DialContext takes precedence over Dial.
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", addr)
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
}

func SADD(key, value string) error {
	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("SADD", key, value)

	return err
}

func SREM(key, value string) error {
	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("SREM", key, value)

	return err
}

func SMEMBERS(key string) ([]string, error) {
	conn := pool.Get()
	defer conn.Close()
	return redis.Strings(conn.Do("SMEMBERS", key))
}



