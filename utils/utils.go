package utils

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/pkg/profile"
	"log"
	"net"
	"net/http"
	"os"
	"runtime/debug"
	"strings"
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

func UmmarshalJson(data string) {

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
