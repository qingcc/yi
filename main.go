package main

import (
	"bufio"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/qingcc/yi/router"
	"io"
	"log"
	"os"
	"time"
)

func ReadLineByLine(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n') //以'\n'为结束符读入一行
		if err != nil || io.EOF == err {
			break
		}
		m[line] = 1
	}

}

var m map[string]int

func main() {
	p := os.Getenv("GOPATH")
	log.Println("p:", p)
	getsha256("")
	go router.InitMetrics(":8082")
	time.Sleep(time.Hour)
}

func getsha256(key string) string {
	url := "/ota/platform/directHotel/getHotelInfos"
	method := "POST"
	data := "{\"uniqueCodes\":[\"8013\"]}"
	arr := method + url + data

	log.Println(arr)
	return ""
}

const (
	message = "POST/ota/platform/directHotel/getHotelInfos{\"uniqueCodes\":[\"8013\"]}"
	secret  = "EWbbJ1jpqe0TvBAk6Bmztk0r0PFmYbkD"
)

func ComputeHmacSha256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	//	fmt.Println(h.Sum(nil))
	sha := hex.EncodeToString(h.Sum(nil))
	fmt.Println(sha)

	hex.EncodeToString(h.Sum(nil))
	return base64.StdEncoding.EncodeToString([]byte(sha))
}

func init() {

	fmt.Println(ComputeHmacSha256(message, secret))
}

func init() {
	secret := secret
	data := message
	fmt.Printf("Secret: %s Data: %s\n", secret, data)

	// Create a new HMAC by defining the hash type and the key (as byte array)
	h := hmac.New(sha256.New, []byte(secret))

	// Write Data to it
	h.Write([]byte(data))

	// Get result and encode as hexadecimal string
	sha := hex.EncodeToString(h.Sum(nil))

	fmt.Println("Result: " + sha)
}
