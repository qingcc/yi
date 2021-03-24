package utils

import (
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// 使用 jsoniter 来json序列化， 效率更高
func ToJson(i interface{}) (js string, err error) {
	var jsBytes []byte
	jsBytes, err = json.Marshal(i)
	return string(jsBytes), err
}

// 使用 jsoniter 来json反序列化， 效率更高
func JsonUnmarshal(jsBytes []byte, str interface{}) (err error) {
	return json.Unmarshal(jsBytes, str)
}

