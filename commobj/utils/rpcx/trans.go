package commobj_trans

type TransRequest struct {
	Query string `json:"q"`     //		text  必填	请求翻译query	    UTF-8编码
	From  string `json:"from"`  //		text  必填	翻译源语言	    语言列表(可设置为auto)
	To    string `json:"to"`    //		TEXT  非必填	译文语言	        语言列表(不可设置为auto)
	Salt  string `json:"salt"`  //		INT	  必填	随机数
	Sign  string `json:"sign"`  //		TEXT  必填	签名	appid+q+salt+密钥 的MD5值
	AppId string `json:"appid"` //      INT   必填	APP ID

	Ssl bool   `json:"-"`
	Url string `json:"-"`
}

type TransAPIRequest struct {
	Query []string `json:"q"`    //		text  必填	请求翻译query	    UTF-8编码
	From  []string `json:"from"` //		text  必填	翻译源语言	    语言列表(可设置为auto)
	To    []string `json:"to"`   //		TEXT  非必填	译文语言	        语言列表(不可设置为auto)
	Salt  []string `json:"salt"` //		INT	  必填	随机数
	Sign  []string `json:"sign"` //		TEXT  必填	签名	appid+q+salt+密钥 的MD5值
	AppId []string `json:"appid"`
}

type TransResponse struct {
	From        string `json:"from"`
	To          string `json:"to"`
	TransResult struct {
		Src uint8  `json:"src"`
		Dst string `json:"dst"`
	} `json:"trans_result"`
}
