package utils

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"runtime/debug"
	"unsafe"
)

func Bytes2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func RecoverPanic(isPrintStack bool) {
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

//func Json2Struct(data string) (stru string) {
//	//Map := make(map[string]interface{})
//	structMap := make(map[string]interface{})
//	if err := json.Unmarshal([]byte(data), &structMap); err != nil {
//		log.Panic("unmarshal failed")
//	}
//	for key, val := range structMap {
//
//	}
//	return
//}
//
//func trans(key string, val interface{}, structMap map[string]interface{}) {
//
//}
//
//func to() {
//
//}

func ToJson(obj interface{}) string {
	bytes, err := json.Marshal(obj)
	if err != nil {

		log.Printf("to json failed (%s), obj %s", err, obj)
	}
	return string(bytes)
}

func UmmarshalJson(data string) {

}
