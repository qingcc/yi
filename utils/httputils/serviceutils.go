package httputils

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"
	"unsafe"
)

const (
	DATA_TYPE_JSON = "json"
	DATA_TYPE_XML = "xml"
	ERROR_DESERIALIZE_XML = 40005
	ERROR_DESERIALIZE_JSON = 40004
	SUCCESS = 40000
	ERROR_REQUEST_TIMEOUT = 40006
	ERROR_SEND_REQUEST = 40007
)

type ServiceHttpClient struct {
	ServiceClient *http.Client
	Timeout time.Duration
}

type HttpRequestContext struct {
	Endpoint string
	ReqBody []byte
	BasicAuth BasicAuthConfig
	IsBasicAuth bool
	Headers map[string]string
	ConcurrentCount int
	DataType string
}

type BasicAuthConfig struct {
	Username string
	Password string
}

func (c *ServiceHttpClient)initHttpClient()  {
	var httpTransport = &http.Transport{
		DialContext:            (&net.Dialer{
			Timeout: 3 * time.Second,
		}).DialContext,
		MaxIdleConns:           20,
		IdleConnTimeout:        10 * time.Second,
		TLSHandshakeTimeout: 2 * time.Second,
	}

	c.ServiceClient = &http.Client{
		Transport:     httpTransport,
	}

	if c.Timeout > 0 * time.Second {
		c.ServiceClient.Timeout = c.Timeout
	}
}

func (c *ServiceHttpClient)SendJson(url string, body []byte, headers map[string]string, concurrentCount int, pRS interface{}) (code int, message string, data []byte) {
	ctx := HttpRequestContext{
		Endpoint:        url,
		ReqBody:         body,
		Headers:         headers,
		DataType: DATA_TYPE_JSON,
		ConcurrentCount: concurrentCount,
	}
	return c.SendWithRequestContext(ctx, http.MethodPost, pRS)
}


func (c *ServiceHttpClient)SendXml(url string, body []byte, headers map[string]string, concurrentCount int, pRS interface{}) (code int, message string, data []byte) {
	ctx := HttpRequestContext{
		Endpoint:        url,
		ReqBody:         body,
		Headers:         headers,
		DataType: DATA_TYPE_XML,
		ConcurrentCount: concurrentCount,
	}
	return c.SendWithRequestContext(ctx, http.MethodPost, pRS)
}
func (c *ServiceHttpClient)SendJsonDelete(url string, headers map[string]string, concurrentCount int, pRS interface{}) (code int, message string, data []byte) {
	ctx := HttpRequestContext{
		Endpoint:        url,
		Headers:         headers,
		DataType: DATA_TYPE_JSON,
		ConcurrentCount: concurrentCount,
	}
	return c.SendWithRequestContext(ctx, http.MethodDelete, pRS)
}


func (c *ServiceHttpClient)GetJson(url string, headers map[string]string, concurrentCount int, pRS interface{}) (code int, message string, data []byte) {
	ctx := HttpRequestContext{
		Endpoint:        url,
		Headers:         headers,
		DataType: DATA_TYPE_JSON,
		ConcurrentCount: concurrentCount,
	}
	return c.SendWithRequestContext(ctx, http.MethodGet, pRS)
}


func (c *ServiceHttpClient)GetXml(url string, headers map[string]string, concurrentCount int, pRS interface{}) (code int, message string, data []byte) {
	ctx := HttpRequestContext{
		Endpoint:        url,
		Headers:         headers,
		DataType: DATA_TYPE_XML,
		ConcurrentCount: concurrentCount,
	}
	return c.SendWithRequestContext(ctx, http.MethodGet, pRS)
}




func (c *ServiceHttpClient)SendWithRequestContext(ctx HttpRequestContext, method string, pRS interface{}) (code int, message string, data []byte)  {
	concurrentCount := 1
	if ctx.ConcurrentCount > 1 {
		concurrentCount = ctx.ConcurrentCount
	}

	if concurrentCount > 1 {
		dataChan := make(chan []byte, concurrentCount)

		var wg sync.WaitGroup

		for i := 0 ; i < concurrentCount ; i++ {
			go func(index int) {
				defer wg.Done()
				if d, err := c.reqData(method, &ctx); err == nil {
					dataChan <- d
				}
			}(i)
		}

		go func() {
			wg.Wait()
			close(dataChan)
		}()

		if c.Timeout <= time.Duration(0) {
			c.Timeout = 10 * time.Minute // set a default timeout
		}

		timer := time.NewTimer(c.Timeout)
		select {
		case respData := <-dataChan:
			code, message = deserializeData(respData, ctx, pRS)
			timer.Stop()
		case t := <- timer.C:
			code = ERROR_REQUEST_TIMEOUT
			message = fmt.Sprint("*** Service timeout at %v", t)
		}
	}else {
		for i := 0; i < 3 ; i++ {
			if d, err := c.reqData(method, &ctx); err == nil {
				data = d
				code, message = deserializeData(d, ctx, pRS)
				break
			} else {
				code = ERROR_SEND_REQUEST
				if strings.Contains(strings.ToLower(err.Error()), "timeout") {
					code = ERROR_REQUEST_TIMEOUT
				}
				message = err.Error()
				switch  {
				case strings.Contains(err.Error(), "TLS handshake timeout"):
				case strings.HasPrefix(err.Error(), "i/o timeout") && strings.Contains(err.Error(), "dial tcp"):
					//do not break, will do retry request
				default:
					break

				}
			}

		}
	}

	if code != SUCCESS {
		log.Printf("[error] request url: %s, error: %s", ctx.Endpoint, message)
		log.Printf("[error] response: %s", Bytes2String(data))
	}
}

func (c *ServiceHttpClient)reqData(method string, ctx *HttpRequestContext) (data []byte, err error) {
	if c.ServiceClient == nil {
		c.initHttpClient()
	}

	if req, newReqErr := http.NewRequest(method, ctx.Endpoint, bytes.NewBuffer(ctx.ReqBody)); newReqErr != nil {
		err = newReqErr
		return
	} else {
		req.Header = make(http.Header)
		for key, val := range ctx.Headers {
			req.Header.Set(key, val)
		}

		if ctx.DataType == DATA_TYPE_JSON {
			req.Header.Set("Accept", "application/json")
		}else if ctx.DataType == DATA_TYPE_XML {
			req.Header.Set("Accept", "application/xml")
		}

		if ctx.IsBasicAuth {
			req.SetBasicAuth(ctx.BasicAuth.Username, ctx.BasicAuth.Password)
		}

		if resp, e := c.ServiceClient.Do(req); e == nil {
			var isGzip bool
			var reader io.ReadCloser
			switch resp.Header.Get("Content-Encoding") {
			case "gzip":
				if reader, err = gzip.NewReader(resp.Body); err != nil {
					log.Println(err)
				}
				isGzip = true
			default:
				reader = resp.Body
			}

			if err == nil && reader != nil {
				if data, err = ioutil.ReadAll(resp.Body); err != nil {
					log.Printf("ioutil.ReadAll() error: %v\n", err)
				}

				if isGzip {
					reader.Close()
				}
			}
			resp.Body.Close()
		}else {
			err = e
			return
		}
	}

	return
}

func deserializeData(d []byte, ctx HttpRequestContext, pRS interface{}) (code int, message string) {
	code = SUCCESS
	message = "Success"
	if pRS != nil {
		if ctx.DataType == DATA_TYPE_XML {
			if err := xml.Unmarshal(d, pRS); err != nil {
				code = ERROR_DESERIALIZE_XML
				message = "invalid xml response error, " + err.Error()
			}
		} else if ctx.DataType == DATA_TYPE_JSON {
			if err := json.Unmarshal(d, pRS); err != nil {
				code = ERROR_DESERIALIZE_JSON
				message = "invalid json response error, " + err.Error()
			}
		}
	}
	return
}


func Bytes2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}