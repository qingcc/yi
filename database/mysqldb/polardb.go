package mysqldb

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"log"
	"sync"
)

var (
	rDBConn *xorm.Engine
	rDBOnce sync.Once
)

func GetRDBConn() *xorm.Engine {
	rDBOnce.Do(func() {
		var err error

		if rDBConn, err = xorm.NewEngine("mysql", "root:123@/test?charset=utf8"); err != nil {
			log.Printf("connection to mysql failed:%s", err.Error())
		} else {
			rDBConn.SetMaxIdleConns(10)
			rDBConn.SetMaxOpenConns(10)
			//名称映射规则主要负责结构体名称到表名和结构体field到表字段的名称映射
			rDBConn.SetTableMapper(core.SnakeMapper{})
		}

	})

	return rDBConn
}
