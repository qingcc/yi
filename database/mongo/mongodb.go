package mongo

import (
	"github.com/globalsign/mgo"
	"log"
	"sync"
)

var (
	connUrl      = "mongodb://47.112.210.86:7017"
	mgoDBSession *mgo.Session
	mgoConnOnce  sync.Once
)

func GetMongoConn() *mgo.Session {
	mgoConnOnce.Do(func() {
		if session, err := mgo.Dial(connUrl); err != nil {
			panic(err)
		} else {
			e := session.DB("test").Login("test", "123456")
			if e != nil {
				log.Println(err.Error())
			}
			mgoDBSession = session
			mgoDBSession.SetMode(mgo.Monotonic, true)
		}
	})
	return mgoDBSession
}
