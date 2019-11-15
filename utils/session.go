package utils

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SetSession(c *gin.Context, key string, val interface{}) {
	sess := sessions.Default(c)
	sess.Set(key, val)
	sess.Save()
}

func GetSession(c *gin.Context, key string) interface{} {
	sess := sessions.Default(c)
	ss := sess.Get(key)
	return ss
}

func HasSession(c *gin.Context, key string) bool {
	sess := sessions.Default(c)
	value := sess.Get(key)
	sess.Save()
	if value == nil {
		return false
	}
	return true
}

func DelSession(c *gin.Context, key string) {
	sess := sessions.Default(c)
	sess.Delete(key)
	sess.Save()
}

func ClearSession(c *gin.Context) {
	sess := sessions.Default(c)
	sess.Clear()
	sess.Save()
}
