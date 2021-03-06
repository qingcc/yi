package redis_utils

import (
	"fmt"
	rediscli "github.com/qingcc/yi/database/redis"
	"log"
	"testing"
	"time"
)

func TestSetKeyNotExistEx(t *testing.T) {
	ok := "r: "
	err := SetKeyNotExistEx("test", 0, "123412341", 5*time.Second)
	log.Println(ok, err)
	err = SetKeyNotExistEx("test", 0, "123412341", 5*time.Second)
	log.Println(ok, err)

	time.Sleep(5 * time.Second)
	err = SetKeyNotExistEx("test", 0, "123412341", 5*time.Second)
	log.Println(ok, err)

	UpdateString2Redis("aaa", 0, "updateString2Redis")
	va, _ := RetrieveStringFromRedis("aaa", 0)
	fmt.Println("val:", va)
}

func TestRetrieveFromRedis(t *testing.T) {
	if res, err := RetrieveStringFromRedis("test_key", rediscli.RedisDBIdx_Common); err == nil {
		log.Println(res)
	} else {
		log.Println("get:", err.Error())
	}
}
