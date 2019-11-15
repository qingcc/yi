package captcha

import (
	"github.com/gin-gonic/gin"
	"github.com/qingcc/goblog/util"
	"strconv"
)

func GetCaptcha(c *gin.Context) {
	width, _ := strconv.Atoi(c.Param("width"))
	height, _ := strconv.Atoi(c.Param("height"))
	d := make([]byte, 4)
	s := NewLen(4)
	char := ""
	d = []byte(s)
	for v := range d {
		d[v] %= 10
		char += strconv.FormatInt(int64(d[v]), 32)
	}
	util.SetSession(c, "captcha_value", char)
	c.Header("Content-Type", "image/png")
	NewImage(d, width, height).WriteTo(c.Writer)
}
func VerifyCaptcha(c *gin.Context, verify_value string) bool {
	value := util.GetSession(c, "captcha_value")
	return value == verify_value
}
