package utils

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"github.com/go-ini/ini"
	"github.com/polaris1119/logger"
	"html/template"
	"log"
	"net"
	"net/smtp"
	"runtime"
	"strconv"
	"strings"
)

type Email struct {
	From     string
	Password string
	Host     string
	Port     string
	To       []string
	Subject  string
	Body     string
	Mailtype string
}

//发送邮件
func InitEmail() Email {
	//读取配置文件
	cfg, err := ini.Load("config/email.ini")
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		logger.Errorf(file+":"+strconv.Itoa(line), "无法打开 'config/email.ini': %v", err)
	}
	key, _ := cfg.Section("email").GetKey("email")
	user := key.Value()
	key, _ = cfg.Section("email").GetKey("host")
	host := key.Value()
	key, _ = cfg.Section("email").GetKey("port")
	port := key.Value()
	key, _ = cfg.Section("email").GetKey("password")
	password := key.Value()
	return Email{From: user, Password: password, Host: host, Port: port}
}

//发送邮件的逻辑函数 type_id: 1 未登录状态, 2 登录状态, 3 更改邮箱
func SendMail(r Email, type_id int) bool {
	hp := strings.Split(r.Host, ":")
	auth = smtp.PlainAuth("", r.From, r.Password, hp[0])

	//code := app.RandInt64(100000, 999999)
	code := int64(666555)
	templateData := struct {
		Name string
		Code int64
	}{
		Name: "code",
		Code: code,
	}
	if err := r.ParseTemplate("views/email/template"+strconv.Itoa(type_id)+".html", templateData); err == nil {
		ok, _ := r.SendEmail()
		str_key := r.To[0] + ":" + strconv.Itoa(type_id)
		err := Set(str_key, code, 60*5)
		if err == false {
			return false
		}
		return ok
	}
	return false
}

var auth smtp.Auth

func (r *Email) SendEmail() (bool, error) {
	header := make(map[string]string)
	header["From"] = "test" + "<" + r.From + ">"
	header["To"] = r.To[0]
	header["Subject"] = r.Subject
	header["Content-Type"] = "text/html; charset=UTF-8"
	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + r.Body
	addr := r.Host + ":" + r.Port
	if err := SendMailUsingTLS(addr, auth, r.From, r.To, []byte(message)); err != nil {
		fmt.Println("send email failed,", err.Error())
		return false, err
	}
	fmt.Println("send email success")
	return true, nil
}

func Dial(addr string) (*smtp.Client, error) {
	conn, err := tls.Dial("tcp", addr, nil)
	if err != nil {
		log.Println("Dialing Error:", err)
		return nil, err
	}
	//分解主机端口字符串
	host, _, _ := net.SplitHostPort(addr)
	return smtp.NewClient(conn, host)
}

//参考net/smtp的func SendMail()
//使用net.Dial连接tls(ssl)端口时,smtp.NewClient()会卡住且不提示err
//len(to)>1时,to[1]开始提示是密送
func SendMailUsingTLS(addr string, auth smtp.Auth, from string,
	to []string, msg []byte) (err error) {
	//create smtp client
	c, err := Dial(addr)
	if err != nil {
		log.Println("Create smpt client error:", err)
		return err
	}
	defer c.Close()
	if auth != nil {
		if ok, _ := c.Extension("AUTH"); ok {
			if err = c.Auth(auth); err != nil {
				log.Println("Error during AUTH", err)
				return err
			}
		}
	}
	if err = c.Mail(from); err != nil {
		return err
	}
	for _, addr := range to {
		if err = c.Rcpt(addr); err != nil {
			return err
		}
	}
	w, err := c.Data()
	if err != nil {
		return err
	}
	_, err = w.Write(msg)
	if err != nil {
		return err
	}
	err = w.Close()
	if err != nil {
		return err
	}
	return c.Quit()
}

//使用自定义模板
func (r *Email) ParseTemplate(templateFileName string, data interface{}) error {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return err
	}
	r.Body = buf.String()
	return nil
}
