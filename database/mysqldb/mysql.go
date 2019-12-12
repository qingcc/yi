package mysqldb

import (
	"database/sql"
	"sync"
)

var(
	dbConn sql.DB
	dbOnce sync.Once
)

func GetConn()  {
	
}
