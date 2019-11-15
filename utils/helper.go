package utils

import (
	"bytes"
	"crypto/md5"
	rand2 "crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"github.com/qingcc/goblog/databases"
	"github.com/shopspring/decimal"
	"html/template"
	"io"
	"math"
	"math/rand"
	"net"
	"reflect"
	"strconv"
	"strings"
	"time"
)

//region 模版自定义函数，可以在模板进行调用
func TemplateFunc() template.FuncMap {
	return template.FuncMap{
		"SizeFormat": SizeFormat,
		"RandAnim": func() string {
			var anim = [5]string{"layui-anim-up", "layui-anim-upbit", "layui-anim-scale", "layui-anim-scaleSpring", "layui-anim-fadein"}
			return anim[rand.Intn(4)]
		},
	}
}

//region Remark: MD5加密 $ Author:Qing
func Strmd5(str string) string {
	w := md5.New()
	w.Write([]byte(str))
	return hex.EncodeToString(w.Sum(nil))
}

//endregion

//生成指定范围的随机数
func RandInt64(min, max int64) int64 {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return rand.Int63n(max-min) + min
}

func GetSjCode(len int64) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes1 := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var i int64
	for i = 0; i < len; i++ {
		//result = append(result, bytes[r.Intn(len(bytes))])  bytes.Count([]byte(str),nil)-1)
		result = append(result, bytes1[r.Intn(bytes.Count(bytes1, nil)-1)])
	}
	return string(result)
}

func CheckErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}

func AddWsLinksToRedis(user_id string) {
	redis_key := "ws_links"
	if res, _ := Exists(redis_key); res {
		value, _ := redis.String(Get(redis_key))
		if value == "" {
			value = user_id
		} else {
			value += "," + user_id
		}
		Set(redis_key, value, -1)
	} else {
		Set(redis_key, user_id, -1)
	}
}

//生成32位随机序列
func RandNewStr(strlen int) string {
	var (
		codes   = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
		codeLen = len(codes)
	)
	data := make([]byte, strlen)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < strlen; i++ {
		idx := rand.Intn(codeLen)
		data[i] = byte(codes[idx])
	}

	return string(data)
}

//region Int64类型的数组去重 去0
func RemoveDuplicateAnd0Int64(list []int64) []int64 {
	var x []int64 = []int64{}
	for _, i := range list {
		if i == 0 {
			continue
		}
		if len(x) == 0 {
			x = append(x, i)
		} else {
			for k, v := range x {
				if i == v {
					break
				}
				if k == len(x)-1 {
					x = append(x, i)
				}
			}
		}
	}
	return x
}

//endregion

//打乱字符串顺序
func Random(strings []string, length int) (string, error) {
	if len(strings) <= 0 {
		return "", errors.New("字符串数组的长度不能为0")
	}

	if length <= 0 || len(strings) < length {
		return "", errors.New("参数length长度不正确")
	}

	for i := len(strings) - 1; i > 0; i-- {
		num := rand.Intn(i + 1)
		strings[i], strings[num] = strings[num], strings[i]
	}

	str := ""
	for i := 0; i < length; i++ {
		str += "," + strings[i]
	}

	return str[1:], nil
}

//字符串数字转为map类型
func Strings2Map(str []string) map[string]int {
	list := make(map[string]int, len(str))
	for key, value := range str {
		list[value] = key + 1
	}
	return list
}

//region Remark: 根据时间戳获取距现在时间 Author; chijian
func UnixTimeTotime(created_time int64) string {
	t := time.Now().Unix() - created_time
	if t < 60 {
		return strconv.FormatInt(t, 64) + "秒前"
	} else if t < 3600 {
		return strconv.FormatInt(int64(math.Floor(float64(t)/60)), 10) + "分钟前"
	} else if t < 86400 {
		return strconv.FormatInt(int64(math.Floor(float64(t)/60/60)), 10) + "小时前"
	} else if t < 604800 {
		return strconv.FormatInt(int64(math.Floor(float64(t)/60/60/24)), 10) + "天前"
	} else if t < 2419200 {
		return strconv.FormatInt(int64(math.Floor(float64(t)/60/60/24/7)), 10) + "周前"
	} else if t < 2592000 {
		return "三周之前"
	} else if t < 31104000 {
		return strconv.FormatInt(int64(math.Floor(float64(t)/60/60/30)), 10) + "月前"
	} else {
		return time.Unix(t, 0).Format("2006-01-02")
	}
}

func GetMd5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func Guid() string {
	b := make([]byte, 48)

	if _, err := io.ReadFull(rand2.Reader, b); err != nil {
		return ""
	}
	return GetMd5(base64.URLEncoding.EncodeToString(b))
}

//region Remark: 获取上个月的开始时间和结束 Author; chijian
func LastMonthStartAndEnd() (time.Time, time.Time) {
	year, month, _ := time.Now().Date()
	thisMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	start := thisMonth.AddDate(0, -1, 0)
	end := thisMonth.AddDate(0, 0, -1)
	return start, end
}

//endregion

//region Remark: 获取该月时间 Author; chijian
func GetThisMonthTime() time.Duration {
	year, month, _ := time.Now().Date()
	start := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	end := start.AddDate(0, 1, 0)
	return end.Sub(start)
}

//endregion

const XRequestedWith = "X-Requested-With"

func IsAjax(c *gin.Context) bool {
	if c.Request.Header.Get(XRequestedWith) == "XMLHttpRequest" {
		return true
	}
	return false
}

func RoundFloat(f float64, m int) float64 {
	floatStr := fmt.Sprintf("%."+strconv.Itoa(m)+"f", f)
	flt, _ := strconv.ParseFloat(floatStr, 64)
	return flt
}

//region Remark: 时间格式 Author; chijian
func GetTimeType(t time.Time, time_type int) time.Time { // t  2009-03-03 15:30:30 +0000 UTC
	switch time_type {
	case 1: //显示年月日 小时 分钟	2009-03-03 15:30:00 +0000 UTC
		return t.Truncate(time.Minute)
	case 2: //显示年月日 小时     2009-03-03 15:00:00 +0000 UTC
		return t.Truncate(time.Hour)
	case 3: //显示年月日			2009-03-03 00:00:00 +0000 UTC
		return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
	case 4: //显示年月				2009-03-01 00:00:00 +0000 UTC
		return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
	case 5: //显示年				2009-01-01 00:00:00 +0000 UTC
		return time.Date(t.Year(), 1, 1, 0, 0, 0, 0, t.Location())
	case 6: //获取本周周一的时间
		return time.Date(t.Year(), t.Month(), t.Day()-int(t.Weekday())+1, 0, 0, 0, 0, time.Local)
	default:
		return t
	}
}

//endregion

//region Remark:生成csv文件并下载 Author:line
func DownCsv(c *gin.Context, fileName string, b *bytes.Buffer) {
	c.Writer.Header().Set("Content-Type", "text/csv")
	c.Writer.Header().Set("Content-Disposition", fmt.Sprintf("attachment;filename=%s", fileName))
	tet := b.String()
	c.String(200, tet)
	c.Next()
}

//endregion

//清空表
func CleanData(table string) {
	gsql := "delete from `" + table + "`" //删除数据
	_, err := databases.Orm.Exec(gsql)
	if err != nil {
		fmt.Println("sql 执行失败::", gsql)
	}
}

//精确计算
func DecimalCal(m float64, n float64, symbol string) float64 {
	M := decimal.NewFromFloat(m)
	N := decimal.NewFromFloat(n)
	ret := float64(0)
	switch symbol {
	case "+":
		ret, _ = M.Add(N).Float64()
		break
	case "-":
		ret, _ = M.Sub(N).Float64()
		break
	case "*":
		ret, _ = M.Mul(N).Float64()
		break
	case "/":
		ret, _ = M.Div(N).Float64()
		break
	}
	return ret
}

//region 获取客户端IP
func ClientIp(c *gin.Context) string {
	remoteAddr := c.Request.RemoteAddr
	if ip := c.Request.Header.Get("XRealIP"); ip != "" {
		remoteAddr = ip
	} else if ip = c.Request.Header.Get("XForwardedFor"); ip != "" {
		remoteAddr = ip
	} else {
		remoteAddr, _, _ = net.SplitHostPort(remoteAddr)
	}

	if remoteAddr == "::1" {
		remoteAddr = "127.0.0.1"
	}

	return remoteAddr
}

//endregion

//region 通过dns服务器8.8.8.8:80获取使用的ip
func PulicIP() string {
	conn, _ := net.Dial("udp", "8.8.8.8:80")
	defer conn.Close()
	localAddr := conn.LocalAddr().String()
	idx := strings.LastIndex(localAddr, ":")
	return localAddr[0:idx]
}

//endregion

//region Remark: 将查询出来的数据结构只取部分字段展示给前端 Author:Qing

//region Remark: 获取指定字段 Author; chijian
// fields "id, order_number, user_info:{id as user_id; mobile; user_name}, bank_card_info:{bank;name;bank_card}")
// 主表字段之间以","隔开, 副表格式 user_info:{...;...;... as ...}, 其中字段之间以";"隔开, 用别名时, 使用 " as "
func Struct2Map(val interface{}, fields string) []interface{} {
	dataByte, _ := json.Marshal(val)
	js, _ := simplejson.NewJson(dataByte)
	ids := strings.Split(fields, ",")
	nodes, err := js.Array()
	if err != nil { //单条数据
		return []interface{}{dealSingleData(val, ids)}
	}
	data := make([]interface{}, len(nodes))

	for key, value := range nodes {
		data[key] = dealSingleData(value, ids)
		//fmt.Println("item:", item)
	}
	//fmt.Println(data)
	return data
}

func dealSingleData(value interface{}, ids []string) map[string]interface{} {
	jsonMap := make(map[string]interface{})
	jsonMap = Json2Map(value)
	item := make(map[string]interface{})
	for _, v := range ids {
		aim, sour := string2key(v)
		for k, _ := range sour {
			s := strings.Replace(sour[k], " ", "", -1)
			a := strings.Replace(aim[k], " ", "", -1)
			item[a] = jsonMap[s]
		}
	}
	return item
}

//region Remark: 对传入的需要取值的字段解析, field是已经用","分隔开的字符串 Author; chijian

//endregion
func string2key(field string) ([]string, []string) {
	//该 field 中是否是要从副表中取字段(格式 user_info:{name; mobile; id as user_id})
	aim := make([]string, 0)  //最后需要展示的key键切片
	sour := make([]string, 0) //源数据中的key键切片
	ids := strings.Split(field, "{")
	if len(ids) == 1 { //主表字段
		return append(aim, field), append(sour, field)
	}
	//需要从副表取字段
	fields := ids[1][:len(ids[1])-1] //具体字段 name; mobile; id as user_id} 需要去除最后的"}"字段
	prefix := ids[0][:len(ids[0])-1] //前缀 ids[0] user_info: 需要去除最后的":"元素

	f_ids := strings.Split(fields, ";") //副表字段以 ";" 隔开

	for _, val := range f_ids { //循环 name; mobile; id as user_id 副表字段
		f_as := strings.Split(val, " as ")
		if len(f_as) == 1 { //副表字段中没有 " as "标签, 副表val字段, 展示的键就是val
			sour = append(sour, prefix+"_"+val)
			aim = append(aim, val)
		} else {
			sour = append(sour, prefix+"_"+f_as[0])
			aim = append(aim, f_as[1]) //副表字段中有 " as "标签, 副表val字段, 展示的键是f_as[1]
		}
	}
	return aim, sour
}

//endregion

//region Remark: 单条数据 json格式转为 map[string]interface{} 格式(其中副表以 table_field 为键存储) Author; chijian
func Json2Map(json_inter interface{}) (jsonMap map[string]interface{}) {
	json_bt, _ := json.Marshal(json_inter)
	//fmt.Println("string:", string(json_bt))
	json.Unmarshal(json_bt, &jsonMap)
	for key, value := range jsonMap {
		if m := reflect.ValueOf(value); m.Kind() == reflect.Map {
			item := make(map[string]interface{})
			v_bt, _ := json.Marshal(value)
			json.Unmarshal(v_bt, &item)
			for k, val := range item {
				jsonMap[key+"_"+k] = val
			}
		}
	}
	return jsonMap
}

//endregion

//endregion
