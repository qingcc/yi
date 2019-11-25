package redisPool

import (
	"github.com/go-redis/redis"
)

func init()  {
	redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // use default Addr
		Password: "",               // no password set
		DB:       0,                // use default DB
	})
}
