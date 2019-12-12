package pgsqldb


import (
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"sync"
	"xorm.io/core"
	"xorm.io/xorm"
)
var(
	rDBConn *xorm.Engine
	rDBOnce sync.Once
)

func GetRDBConn() *xorm.Engine {
	rDBOnce.Do(func() {
		var err error
		if rDBConn, err = xorm.NewEngine("postgres", fmt.Sprintf("host=127.0.0.1 port=5432 user=user password=password dbname=go sslmode=disable")); err != nil {
			log.Printf("connection to mysql failed:%s", err.Error())
		}else {
			rDBConn.SetMaxIdleConns(10)
			rDBConn.SetMaxOpenConns(10)
			//名称映射规则主要负责结构体名称到表名和结构体field到表字段的名称映射
			rDBConn.SetTableMapper(core.SnakeMapper{})
		}
	})
	return rDBConn
}
