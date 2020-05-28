package utils

import (
	"bufio"
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/pkg/profile"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"reflect"
	"runtime/debug"
	"strings"
	"unicode"
	"unsafe"
)

func Bytes2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func RecoverFromPanic(isPrintStack bool) {
	if r := recover(); r != nil {
		log.Println(r)
		if isPrintStack {
			debug.PrintStack()
		}
	}
	return
}

//region Author: qing  标准输入
func Stdin(n int) string {
	if n == 0 {
		n = 1024
	}
	var buf = make([]byte, 0, n)
	if m, err := os.Stdin.Read(buf); err == nil {
		return string(buf[:m])
	} else {
		log.Printf("read error：%s", err.Error())
	}
	return ""
}

var (
	firstname, lastname string
	age                 int
)

func Scanf() {
	fmt.Println("please input firstname, lastname, age:")
	fmt.Scanln(&firstname, &lastname, &age)
	fmt.Println("firstname:", firstname, "lastname:", lastname, "age:", age)
}

func Stdin2() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("please enter something:")
	input, err := inputReader.ReadString('\n')
	if err == nil {
		fmt.Println("the input was:", input)
	}
}

//endregion

func ToJson(obj interface{}) string {
	bytes, err := json.Marshal(obj)
	if err != nil {
		log.Printf("to json failed (%s), obj %s", err, obj)
	}
	return string(bytes)
}

////region Remark: 将查询出来的数据结构只取部分字段展示给前端 Author:Qing
//
////region Remark: 获取指定字段 Author; chijian
//// fields "id, order_number, user_info:{id as user_id; mobile; user_name}, bank_card_info:{bank;name;bank_card}")
//// 主表字段之间以","隔开, 副表格式 user_info:{...;...;... as ...}, 其中字段之间以";"隔开, 用别名时, 使用 " as "
//func Struct2Map(val interface{}, fields string) []interface{} {
//	dataByte, _ := json.Marshal(val)
//	js, _ := simplejson.NewJson(dataByte)
//	ids := strings.Split(fields, ",")
//	nodes, err := js.Array()
//	if err != nil { //单条数据
//		return []interface{}{dealSingleData(val, ids)}
//	}
//	data := make([]interface{}, len(nodes))
//
//	for key, value := range nodes {
//		data[key] = dealSingleData(value, ids)
//		//fmt.Println("item:", item)
//	}
//	//fmt.Println(data)
//	return data
//}
//
//func dealSingleData(value interface{}, ids []string) map[string]interface{} {
//	jsonMap := make(map[string]interface{})
//	jsonMap = Json2Map(value)
//	item := make(map[string]interface{})
//	for _, v := range ids {
//		aim, sour := string2key(v)
//		for k, _ := range sour {
//			s := strings.Replace(sour[k], " ", "", -1)
//			a := strings.Replace(aim[k], " ", "", -1)
//			item[a] = jsonMap[s]
//		}
//	}
//	return item
//}
//
////region Remark: 对传入的需要取值的字段解析, field是已经用","分隔开的字符串 Author; chijian
//
////endregion
//func string2key(field string) ([]string, []string) {
//	//该 field 中是否是要从副表中取字段(格式 user_info:{name; mobile; id as user_id})
//	aim := make([]string, 0)  //最后需要展示的key键切片
//	sour := make([]string, 0) //源数据中的key键切片
//	ids := strings.Split(field, "{")
//	if len(ids) == 1 { //主表字段
//		return append(aim, field), append(sour, field)
//	}
//	//需要从副表取字段
//	fields := ids[1][:len(ids[1])-1] //具体字段 name; mobile; id as user_id} 需要去除最后的"}"字段
//	prefix := ids[0][:len(ids[0])-1] //前缀 ids[0] user_info: 需要去除最后的":"元素
//
//	f_ids := strings.Split(fields, ";") //副表字段以 ";" 隔开
//
//	for _, val := range f_ids { //循环 name; mobile; id as user_id 副表字段
//		f_as := strings.Split(val, " as ")
//		if len(f_as) == 1 { //副表字段中没有 " as "标签, 副表val字段, 展示的键就是val
//			sour = append(sour, prefix+"_"+val)
//			aim = append(aim, val)
//		} else {
//			sour = append(sour, prefix+"_"+f_as[0])
//			aim = append(aim, f_as[1]) //副表字段中有 " as "标签, 副表val字段, 展示的键是f_as[1]
//		}
//	}
//	return aim, sour
//}
//
////endregion
//
////region Remark: 单条数据 json格式转为 map[string]interface{} 格式(其中副表以 table_field 为键存储) Author; chijian
//func Json2Map(json_inter interface{}) (jsonMap map[string]interface{}) {
//	json_bt, _ := json.Marshal(json_inter)
//	//fmt.Println("string:", string(json_bt))
//	json.Unmarshal(json_bt, &jsonMap)
//	for key, value := range jsonMap {
//		if m := reflect.ValueOf(value); m.Kind() == reflect.Map {
//			item := make(map[string]interface{})
//			v_bt, _ := json.Marshal(value)
//			json.Unmarshal(v_bt, &item)
//			for k, val := range item {
//				jsonMap[key+"_"+k] = val
//			}
//		}
//	}
//	return jsonMap
//}
//
////endregion
//
////endregion

func Struct2Map(obj interface{}) map[string]interface{} {
	typeOf := reflect.TypeOf(obj)
	value := reflect.ValueOf(obj)

	var result = make(map[string]interface{})
	for i := 0; i < typeOf.NumField(); i++ {
		if typeOf.Field(i).Anonymous {
			for k, v := range Struct2Map(value.Field(i).Interface()) {
				result[k] = v
			}
			continue
		}

		name := typeOf.Field(i).Tag.Get("db")
		if name == "" {
			name = typeOf.Field(i).Tag.Get("json")
		}

		if name == "" {
			name = typeOf.Field(i).Name
		}
		result[name] = value.Field(i).Interface()
	}
	return result
}

func String2Bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}

	return *(*[]byte)(unsafe.Pointer(&h))
}

var ErrNoNodeNameEnvSet = errors.New("no env settings by 'MY_NODE_NAME'")

func GetK8SNodeIP() (ipAddr string, err error) {
	nodeName := os.Getenv("MY_NODE_NAME")
	if nodeName == "" {
		err = ErrNoNodeNameEnvSet
		return
	}

	ipAddr = strings.TrimPrefix(nodeName, "cn-shenzhen.")
	if ip := net.ParseIP(ipAddr); ip == nil {
		err = fmt.Errorf("invalid ip address parse by node name: %s", nodeName)
		ipAddr = ""
	}

	if err != nil {
		log.Println(err)
	}

	return
}

func GetInternalIp() (string, error) {
	var addrStr string
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Printf("get IP address failed: %v", err)
	} else {

		for _, addr := range addrs {
			if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					addrStr = ipnet.IP.String()
					if strings.HasPrefix(addrStr, "10.") ||
						strings.HasPrefix(addrStr, "172.") ||
						strings.HasPrefix(addrStr, "192.") {

						log.Println(addrStr)
						return addrStr, nil
					}
				}
			}
		}
	}

	return "", errors.New("no expected internal ip found")
}

//region Remark: pprof监测 Author:qing
func Pprof(addr string) {
	if addr == "" {
		addr = ":6061"
	}
	go func() {
		log.Println(http.ListenAndServe(addr, nil))
	}()
}

//endregion

func Pprof2File(f func()) {
	stopper := profile.Start(profile.CPUProfile, profile.ProfilePath(".")) // 开始性能分析, 返回一个停止接口
	defer stopper.Stop()                                                   // 在被测试程序结束时停止性能分析
	f()
}

func Strmd5(str string) string {
	hash := md5.New()
	io.WriteString(hash, str) //RatePlanCode房型下唯一
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func HasChinese(str string) bool {
	if str == "" {
		return true
	}
	for _, v := range str {
		if unicode.Is(unicode.Han, v) {
			return true
		}
	}
	return false
}
