package utils

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	"io/ioutil"
	"net/http"
	"net/url"
)

/*
*使用百度翻译api接口,翻译文本
*
 */

var (
	Url  = "http://api.fanyi.baidu.com/api/trans/vip/translate"  //http访问翻译
	Urls = "https://fanyi-api.baidu.com/api/trans/vip/translate" //https访问翻译

	appId      = "20190621000309439"    //APP ID
	secryptkey = "3gFJdY73x_NYKAMUU0UB" //密钥
)

func Trans(words, to string) string { //翻译函数
	body := translate(words, to)
	js, err := simplejson.NewJson(body)
	if err != nil {
		panic(err.Error())
	}
	dst := js.Get("trans_result").GetIndex(0).Get("dst").MustString()

	if dst == "" {
		fmt.Println("error", string(body))
	}
	return dst
}

func translate(query, to string) []byte {
	salt := "3"
	resp, err := http.PostForm(Url, url.Values{
		"q":     {query},                                     //		text  必填	请求翻译query	    UTF-8编码
		"from":  {"zh"},                                      //		text  必填	翻译源语言	    语言列表(可设置为auto)
		"to":    {to},                                        //		TEXT  非必填	译文语言	        语言列表(不可设置为auto)
		"appid": {appId},                                     //        INT   必填	APP ID
		"salt":  {salt},                                      //		INT	  必填	随机数
		"sign":  {Strmd5(appId + query + salt + secryptkey)}, //		TEXT  必填	签名	appid+q+salt+密钥 的MD5值
	})
	if err != nil {
		fmt.Println("请求出错, 错误信息:", err.Error())
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	fmt.Println("body:", string(body))
	return body
}

//region MD5加密
