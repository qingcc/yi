package redis_utils

import (
	"log"
	"testing"
	"time"
)

func TestSetKeyNotExistEx(t *testing.T) {
	ok := "r: "
	err := SetKeyNotExistEx("test", 0, "123412341", 5 * time.Second)
	log.Println(ok, err)
	err = SetKeyNotExistEx("test", 0, "123412341", 5 * time.Second)
	log.Println(ok, err)

	time.Sleep(5 * time.Second)
	err = SetKeyNotExistEx("test", 0, "123412341", 5 * time.Second)
	log.Println(ok, err)
}
