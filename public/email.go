package public

import (
	"bytes"
	"github.com/jordan-wright/email"
	"html/template"
	"log"
	"net/smtp"
	"strings"
)

type EmailConfig struct {
	id   string
	host string
	port string
	user string
	pass string
}

var config *EmailConfig

func init() {
	DB.Raw("select * from email where id = 1").Scan(config)
}

func SendEmail(receiver string) {
	auth := smtp.PlainAuth("", config.user, config.pass, config.host)
	msg := []byte("Subject: 这里是标题内容\r\n\r\n" + "这里是正文内容\r\n")
	err := smtp.SendMail(config.host+config.port, auth, config.user, []string{receiver}, msg)
	if err != nil {
		log.Fatal("failed to send EmailConfig:", err)
	}
}

///SendHTMLEmail: 发送HTML格式邮件
func SendHTMLEmail(receiver string, path string, variables interface{}) error {
	e := email.NewEmail()
	e.Subject = "Micah在线视频后台管理"
	e.To = strings.Split(receiver, ",")
	t := template.Must(template.ParseFiles(path))
	body := new(bytes.Buffer)
	//作为变量传递给html模板
	t.Execute(body, variables)
	e.HTML = body.Bytes()
	auth := smtp.PlainAuth("", config.user, config.pass, config.host)
	err := e.Send(config.host+config.port, auth)
	log.Fatal("failed to send email:", err)
	return err
}
