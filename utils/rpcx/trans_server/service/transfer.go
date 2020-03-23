package service

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/qingcc/yi/commobj"
	trans_service "github.com/qingcc/yi/commobj/utils"
	commobj_trans "github.com/qingcc/yi/commobj/utils/rpcx"
	"github.com/qingcc/yi/utils"
	"github.com/qingcc/yi/utils/httputils"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

/*
*使用百度翻译api接口,翻译文本
*
 */

var (
	Url  = "http://api.fanyi.baidu.com/api/trans/vip/translate"  //http访问翻译
	Urls = "https://fanyi-api.baidu.com/api/trans/vip/translate" //https访问翻译

	appId       = "20190621000309439"    //APP ID
	secryptkey  = "3gFJdY73x_NYKAMUU0UB" //密钥
	httpService = httputils.ServiceHttpClient{Timeout: 30 * time.Second}
	header      = map[string]string{"Content-Type": "application/x-www-form-urlencoded", "Accept": "application/json"}
)

func getSign(req trans_service.Args) (domain string, data []byte) {
	salt := strconv.Itoa(rand.Intn(10000))
	u := url.Values{
		"q":     {req.Query},                                           //		text  必填	请求翻译query	    UTF-8编码
		"from":  {req.From},                                            //		text  必填	翻译源语言	    语言列表(可设置为auto)
		"to":    {req.To},                                              //		TEXT  非必填	译文语言	        语言列表(不可设置为auto)
		"appid": {appId},                                               //        INT   必填	APP ID
		"salt":  {salt},                                                //		INT	  必填	随机数
		"sign":  {utils.Strmd5(appId + req.Query + salt + secryptkey)}, //		TEXT  必填	签名	appid+q+salt+密钥 的MD5值
	}
	data = []byte(u.Encode())
	if req.Ssl {
		domain = Urls
	} else {
		domain = Url
	}
	return
}

//old
func Trans(words, from, to string, ssl bool) string { //翻译函数
	if from == "" {
		from = "zh"
	}
	body := translate(words, from, to)
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

func Transfer(transReq trans_service.Args, reply *commobj_trans.TransResponse) {
	domain, data := getSign(transReq)
	SendTransData(domain, data, reply, true)
	return
}

func SendTransData(domain string, data []byte, res interface{}, storage bool) (httpMessageResult commobj.HttpMessageResult) {
	httpMessageResult.Start = time.Now()
	resultMessageArr := make([]string, 0)
	httpMessageResult.Req = string(data)
	resultMessageArr = append(resultMessageArr, "[Marshal Request]Success")
	code, message, data := httpService.SendJson(domain, data, header, 1, res)
	httpMessageResult.ResultCode = code
	httpMessageResult.SplResultCode = code
	httpMessageResult.Url = domain
	httpMessageResult.Res = string(data)
	if storage {
		log.Printf("请求对象==>%v\n", httpMessageResult.Req)
		log.Printf("返回对象==>%v\n", httpMessageResult.Res)
	}

	if code == commobj.SUCCESS {
		httpMessageResult.Success = true
		resultMessageArr = append(resultMessageArr, "[SendJson Message]Success")
	} else {
		resultMessageArr = append(resultMessageArr, "[SendJson Message]Failed "+message)
	}
	httpMessageResult.ResultMessage = strings.Join(resultMessageArr, ",")
	httpMessageResult.SplResultMessage = strings.Join(resultMessageArr, ",")
	httpMessageResult.End = time.Now()
	return
}

func translate(query, from, to string) []byte {
	salt := strconv.Itoa(rand.Intn(10000))
	resp, err := http.PostForm(Url, url.Values{
		"q":     {query},                                           //		text  必填	请求翻译query	    UTF-8编码
		"from":  {from},                                            //		text  必填	翻译源语言	    语言列表(可设置为auto)
		"to":    {to},                                              //		TEXT  非必填	译文语言	        语言列表(不可设置为auto)
		"appid": {appId},                                           //        INT   必填	APP ID
		"salt":  {salt},                                            //		INT	  必填	随机数
		"sign":  {utils.Strmd5(appId + query + salt + secryptkey)}, //		TEXT  必填	签名	appid+q+salt+密钥 的MD5值
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
