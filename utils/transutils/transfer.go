package transutils

import (
	"encoding/json"
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/qingcc/yi/commobj"
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

	appId      		= "20190621000309439"    //APP ID
	secryptkey 		= "3gFJdY73x_NYKAMUU0UB" //密钥
	httpService 	= httputils.ServiceHttpClient{Timeout:30 * time.Second}
	header          = map[string]string{"Content-Type": "application/x-www-form-urlencoded", "Accept": "application/json",
		"Accept-encoding": "gzip, deflate"}
)
func (t *TransRequest)getSign() {
	t.Sign = utils.Strmd5(t.AppId + t.Query + t.Salt + secryptkey)
	if t.Ssl {
		t.Url = Urls
	}else {
		t.Url = Url
	}
	return
}


type TransRequest struct {
	Query string `json:"q"`         //		text  必填	请求翻译query	    UTF-8编码
	From string `json:"from"`       //		text  必填	翻译源语言	    语言列表(可设置为auto)
	To string `json:"to"`           //		TEXT  非必填	译文语言	        语言列表(不可设置为auto)
	Salt string `json:"salt"`       //		INT	  必填	随机数
	Sign string `json:"sign"`		//		TEXT  必填	签名	appid+q+salt+密钥 的MD5值
	AppId string `json:"appid"`     //      INT   必填	APP ID

	Ssl bool `json:"-"`
	Url string `json:"-"`
}

type TransResponse struct {

}

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



func Transfer(query, from, to string, ssl bool) (res TransResponse) {
	req := TransRequest{
		Query: query,
		From:  from,
		To:    to,
		Salt:  strconv.Itoa(rand.Intn(10000)),
		Ssl:   ssl,
		AppId:appId,
	}
	req.getSign()
	SendTransData(req, res, true)
	return
}


func SendTransData(req TransRequest, res interface{}, storage bool) (httpMessageResult commobj.HttpMessageResult) {
	httpMessageResult.Start = time.Now()
	resultMessageArr := make([]string, 0)

	reqData, err := json.Marshal(req)
	if err != nil {
		resultMessageArr = append(resultMessageArr, "[Marshal Request]fail"+err.Error())
	}
	httpMessageResult.Req = string(reqData)
	resultMessageArr = append(resultMessageArr, "[Marshal Request]Success")
	code, message, data := httpService.SendJson(req.Url, reqData, header, 1, res)
	httpMessageResult.ResultCode = code
	httpMessageResult.SplResultCode = code
	httpMessageResult.Url = req.Url
	httpMessageResult.Res = string(data)
	if storage {
		log.Printf("请求对象==>%v\n", httpMessageResult.Req)
		log.Printf("返回对象==>%v\n", httpMessageResult.Res)
	}

	if code == commobj.SUCCESS {
		httpMessageResult.Success = true
		resultMessageArr = append(resultMessageArr, "[SendJson Message]Success")
	}else {
		resultMessageArr = append(resultMessageArr, "[SendJson Message]Failed " + message)
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
