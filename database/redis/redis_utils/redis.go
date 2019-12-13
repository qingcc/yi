package redis_utils

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	rediscli "github.com/qingcc/yi/database/redis"
	"go.uber.org/zap"
	"log"
	"time"
)

type funcGetFromDB func() (interface{}, error)

func RetrieveOrSetJsonFromRedis(key string, dbIdx int, val interface{}, getFromDB funcGetFromDB) (err error) {
	c := rediscli.GetConn()
	defer c.Close()

	toGetFromDB := false
	if err = c.Send("SELECT", dbIdx); err == nil {
		if data, e := redis.Bytes(c.Do("GET", key)); e != nil {
			log.Println("==== Can not get the data from redis", zap.Error(e))
			toGetFromDB = true
		}else {
			if err = json.Unmarshal(data, val); err != nil {
				log.Print("failed to unmashal data"+string(data), err)
				toGetFromDB = true
			}
		}
	}else {
		log.Printf(fmt.Sprintf("==== Failed to select DB, DB Index: %d", dbIdx))
	}

	if toGetFromDB {
		log.Println("-----Get data from database and set to redis")

		var v interface{}
		if v, err = getFromDB(); err != nil {
			log.Println("failed to get data from database", zap.Error(err))
		}else {
			if jsonData, e := json.Marshal(v); e != nil {
				log.Println(fmt.Sprintf("marshal data failed %+v", e), zap.Error(e))
			}else {
				if e := c.Send("SET", key, jsonData); e == nil {
					if e := c.Flush(); e != nil {
						log.Println("flush commands failed:", zap.NamedError("error", e))
					}
				} else {
					log.Println("Send set commands failed", zap.NamedError("error", e))
				}

				if e := json.Unmarshal(jsonData, val); e != nil {
					log.Println("unmashal data failed", zap.NamedError("error", e))
				}
			}
		}
	}
	return
}

//设置key， value
func Update2Redis(key string, dbIdx int, val string) (err error) {
	c := rediscli.GetConn()
	defer c.Close()

	if jsonData, e := json.Marshal(val); e != nil {
		err = e
		log.Printf(fmt.Sprintf("marshal data failed %+v", val), e.Error())
	}else {
		if err = c.Send("SELECT", dbIdx); err == nil {
			//SET resource_name my_random_value NX PX 30000
			if err = c.Send("SET", key, jsonData); err == nil {
				if err = c.Flush(); err != nil {
					log.Println("flush commands failed:", zap.Error(err))
				}
			}else {
				log.Printf("Send set commands failed")
			}
		} else {
			log.Println("Select redis database failed", err)
		}
	}

	return
}


func UpdateStringToRedisEx(key string, dbIdx int, val string, timeout time.Duration) (err error) {
	c := rediscli.GetConn()
	defer c.Close()

	timeoutSec := int(timeout / time.Second)
	if err = c.Send("SELECT", dbIdx); err == nil {
		if err = c.Send("SETEX", key, timeoutSec, val); err == nil {
			if err = c.Flush(); err != nil {
				log.Println("flush commands failed:", zap.Error(err))
			}
		} else {
			log.Println("Send setex commands failed", zap.Error(err))
		}
	} else {
		log.Println("Select redis database failed", zap.Error(err))
	}

	return
}

func UpdateToRedisEx(key string, dbIdx int, val interface{}, timeout time.Duration) (err error) {
	c := rediscli.GetConn()
	defer c.Close()
	// set to redis

	timeoutSec := int(timeout / time.Second)
	if jsonData, e := json.Marshal(val); e != nil {
		err = e
		log.Println(fmt.Sprintf("marshal data failed %+v", val), zap.Error(e))
	} else {
		if err = c.Send("SELECT", dbIdx); err == nil {
			if _, err = c.Do("SETEX", key, timeoutSec, jsonData); err != nil {
				log.Println("SETEX commands failed:", err)
			}
		} else {
			log.Println("Select redis database failed", err)
		}
	}

	return
}


//设置key, 若key值不存在，设置key，val，返回OK。key存在，返回nil。
func SetKeyNotExistEx(key string, dbIdx int, val string, timeout time.Duration) (err error) {
	c := rediscli.GetConn()
	defer c.Close()

	timeoutSec := int(timeout / time.Millisecond)
	if err = c.Send("SELECT", dbIdx); err == nil {
		//SET resource_name my_random_value NX PX 30000
		if _, err = redis.String(c.Do("SET", key, val, "NX", "PX", timeoutSec)); err != nil {
			if err != redis.ErrNil {
				log.Printf("redis command failed, %v: SET %s %s NX PX %d", err, key, val, timeoutSec)
			}
		}
	} else {
		log.Printf("*** Failed to select DB, DB Index: %d", dbIdx)
	}
	return
}

func RetrieveFromRedis(key string, dbIdx int, val interface{}) (err error) {
	c := rediscli.GetConn()
	defer c.Close()

	if err = c.Send("SELECT", dbIdx); err == nil {
		if data, e := redis.Bytes(c.Do("GET", key)); e != nil {
			if e == redis.ErrNil {
				// no data return
			} else {
				log.Printf("*** Can not get the data from redis: %s, key: %s, db index: %d", e, key, dbIdx)
			}
			err = e
		}else {
			if err = json.Unmarshal(data, val); err != nil {
				log.Println("*** failed to unmarshal data:"+string(data), err)
			}
		}
	}else {
		log.Printf("*** Failed to select DB, DB Index: %d", dbIdx)
	}
	return
}
func RetrieveStringFromRedis(key string, dbIdx int, val interface{}) (err error) {
	c := rediscli.GetConn()
	defer c.Close()

	if err = c.Send("SELECT", dbIdx); err == nil {
		if data, e := redis.Bytes(c.Do("GET", key)); e != nil {
			if e == redis.ErrNil {
				// no data return
			} else {
				log.Printf("*** Can not get the data from redis: %s, key: %s, db index: %d", e, key, dbIdx)
			}
			err = e
		}else {
			val = string(data)
		}
	}else {
		log.Printf("*** Failed to select DB, DB Index: %d", dbIdx)
	}
	return
}


func GetKeysWithPattern(dbIndex int, pattern string) ([]string, error) {
	c := rediscli.GetConn()
	defer c.Close()

	var err error
	keys := make([]string, 0)
	if err = c.Send("SELECT", dbIndex); err == nil {
		if byteSlices, e := redis.ByteSlices(c.Do("KEYS", pattern+"*")); e != nil {
			log.Println("*** Can not get the data from redis")
			err = e
		} else {
			for _, val := range byteSlices {
				keys = append(keys, string(val))
			}
		}
	} else {
		log.Printf("==== Failed to select DB, DB Index: %d", dbIndex)
	}

	return keys, nil
}


func GetKeyMapWithPattern(dbIndex int, pattern string) (map[string]string, error) {
	keysMap := make(map[string]string)
	if strList, err := GetKeysWithPattern(dbIndex, pattern); err == nil {
		for _, str := range strList {
			keysMap[str] = str
		}
	} else {
		return keysMap, err
	}

	return keysMap, nil
}

func UpdateStringToRedis(key string, dbIdx int, val string) (err error) {
	c := rediscli.GetConn()
	defer c.Close()

	if err = c.Send("SELECT", dbIdx); err == nil {
		if err = c.Send("SET", key, val); err == nil {
			if err = c.Flush(); err != nil {
				log.Println("flush commands failed:", zap.Error(err))
			}
		} else {
			log.Println("Send set commands failed", zap.Error(err))
		}
	} else {
		log.Println("Select redis database failed", zap.Error(err))
	}

	return
}

func RemoveKeysFromRedis(keys []string, dbIdx int) (err error) {
	c := rediscli.GetConn()
	defer c.Close()
	// remove from redis
	if err = c.Send("SELECT", dbIdx); err == nil {
		args := make([]interface{}, len(keys))
		for i := range keys {
			args[i] = keys[i]
		}
		if err = c.Send("DEL", args...); err == nil {
			if err = c.Flush(); err != nil {
				log.Println(" *** [error] flush commands failed:", err)
			}
		} else {
			log.Println("*** [error] Send set commands failed", err)
		}
	} else {
		log.Println("*** [error] Select redis database failed", err)
	}
	return
}

func ExistsKey(key string, dbIdx int) (exist bool) {
	c := rediscli.GetConn()
	defer c.Close()

	if err := c.Send("SELECT", dbIdx); err == nil {
		if flag, err := redis.Int(c.Do("EXISTS", key)); err == nil {

			if flag == 1 {
				exist = true
			}
		} else {
			log.Println("[error] **** EXISTS commands failed:", err)
		}
	} else {
		log.Println("[error] ****  Select redis database failed", err)
	}

	return
}