package utils

import (
	"log"
	"runtime/debug"
	"unsafe"
)

func Bytes2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func RecoverPanic(isPrintStack bool)  {
	if r := recover(); r != nil {
		log.Println(r)
		if isPrintStack {
			debug.PrintStack()
		}
	}
	return
}
