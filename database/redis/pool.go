package rediscli

import (
	"github.com/gomodule/redigo/redis"
	"log"
	"sync"
	"time"
)

var(
	redisPool *redis.Pool
	redisOnce sync.Once
)

func GetConn() redis.Conn {
	return redisPool.Get()
}

func init() {
	redisOnce.Do(func() {
	log.Printf("init a new redis pool")
	redisPool = newPool()
	})
}

func newPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:         20,
		MaxActive:       10,
		IdleTimeout:     15 * time.Second,
		Wait:            true,
		Dial: func() (conn redis.Conn, e error) {
			dialPasswordOption := redis.DialPassword("123456")
			c, err := redis.Dial("tcp", "47.112.210.86:6379", dialPasswordOption)
			if err != nil {
				log.Println("dial to redis failed")
				return c, err
			}
			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}