package mysqldb

import (
	"database/sql"
	"fmt"
	"github.com/gocraft/dbr"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

var (
	mysqlConn *dbr.Connection
	mysqlOnce sync.Once
)

func GetMysqlConn() *dbr.Connection {
	mysqlOnce.Do(func() {
		user := "user"
		pwd := "password"
		db := "mysqldb"
		host := "127.0.0.1"
		dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true&loc=Local&multiStatements=true",
			user, pwd, host, db)

		var err error
		if mysqlConn, err = dbr.Open("mysql", dsn, nil); err != nil {
			log.Panicln(err)
		} else {
			maxIdle, maxOpen := getDBConnectionCounts()
			mysqlConn.SetMaxIdleConns(maxIdle)
			mysqlConn.SetMaxOpenConns(maxOpen)
			mysqlConn.SetConnMaxLifetime(10 * 60 * time.Second)
		}
	})
	return mysqlConn
}

func getDBConnectionCounts() (maxIdle, maxOpen int) {
	maxIdle = 2
	maxOpen = 10
	if countVal := os.Getenv("MYSQL_IDLE_CONN_COUNT"); countVal != "" {
		if c, atoiErr := strconv.Atoi(countVal); atoiErr == nil {
			maxIdle = c
		} else {
			log.Println("[error] parse MYSQL_IDLE_CONN_COUNT error, ", atoiErr)
		}
	}

	if countVal := os.Getenv("MYSQL_OPEN_CONN_COUNT"); countVal != "" {
		if c, atoiErr := strconv.Atoi(countVal); atoiErr == nil {
			maxOpen = c
		} else {
			log.Println("[error] parse MYSQL_OPEN_CONN_COUNT error, ", atoiErr)
		}
	}

	log.Printf("[info] init mysql connection pool, max open %d, max idle %d", maxOpen, maxIdle)

	return
}
