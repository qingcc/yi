module github.com/qingcc/yi

go 1.13

require (
	github.com/bitly/go-simplejson v0.5.0
	github.com/bmizerany/assert v0.0.0-20160611221934-b7ed37b82869 // indirect
	github.com/boombuler/barcode v1.0.0
	github.com/gin-gonic/gin v1.4.0
	github.com/globalsign/mgo v0.0.0-20181015135952-eeefdecb41b8
	github.com/go-ini/ini v1.51.0
	github.com/go-sql-driver/mysql v1.4.1

	github.com/go-xorm/core v0.6.3 // indirect
	github.com/gocraft/dbr v0.0.0-20190714181702-8114670a83bd // indirect
	github.com/gofrs/uuid v3.2.0+incompatible // indirect
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/jackc/pgx v3.6.0+incompatible // indirect
	github.com/lib/pq v1.1.1
	github.com/polaris1119/logger v0.0.0-20170422061149-0233d014769e
	github.com/qingcc/goblog v0.0.0-20191115095047-0b255681b3d9
	github.com/robfig/cron v1.2.0 // indirect
	github.com/samuel/go-zookeeper v0.0.0-20180130194729-c4fab1ac1bec
	github.com/shopspring/decimal v0.0.0-20180709203117-cd690d0c9e24
	github.com/siddontang/go v0.0.0-20180604090527-bdc77568d726
	github.com/smallnest/rpcx v0.0.0-20191202025149-2fd1f4f7e90c
	github.com/tealeg/xlsx v1.0.5
	github.com/typa01/go-utils v0.0.0-20181126045345-a86b05b01c1e
	go.uber.org/atomic v1.5.1 // indirect
	go.uber.org/multierr v1.4.0 // indirect
	go.uber.org/ratelimit v0.1.0
	go.uber.org/zap v1.13.0
	golang.org/x/lint v0.0.0-20191125180803-fdd1cda4f05f // indirect
	golang.org/x/net v0.0.0-20191112182307-2180aed22343 // indirect
	golang.org/x/tools v0.0.0-20191210221141-98df12377212 // indirect
	golang.org/x/vgo v0.0.0-20180912184537-9d567625acf4 // indirect
	google.golang.org/appengine v1.6.5 // indirect
)

replace (
	github.com/go-xorm/builder => xorm.io/builder v0.3.6
	github.com/go-xorm/core => xorm.io/core v0.7.2-0.20190928055935-90aeac8d08eb
)
