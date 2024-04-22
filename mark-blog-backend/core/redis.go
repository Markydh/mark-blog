package core

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

var pool *redis.Pool

func InitRedis()(*redis.Pool){
    pool = &redis.Pool{
        Dial: func() (redis.Conn, error) {
            return redis.Dial("tcp", "localhost:6379")
        },
        TestOnBorrow: func(c redis.Conn, t time.Time) error {
            _, err := c.Do("PING")
            return err
        },
    }
	return pool
}
