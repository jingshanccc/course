package public

import (
	"bytes"
	"context"
	"gitee.com/jingshanccc/course/public/middleware/redis"
	"gitee.com/jingshanccc/course/public/util"
	"github.com/jordan-wright/email"
	"html/template"
	"log"
	"net/smtp"
	"strings"
	"time"
)

type EmailConfig struct {
	Id   int
	Host string
	Port string
	User string
	Pass string
}

func SendEmail(receiver string) {
	var emailConfig EmailConfig
	DB.Raw("select * from email where id = 1 limit 1").Find(&emailConfig)
	auth := smtp.PlainAuth("", emailConfig.User, emailConfig.Pass, emailConfig.Host)
	msg := []byte("Subject: 这里是标题内容\r\n\r\n" + "这里是正文内容\r\n")
	err := smtp.SendMail(emailConfig.Host+emailConfig.Port, auth, emailConfig.User, []string{receiver}, msg)
	if err != nil {
		log.Fatal("failed to send EmailConfig:", err)
	}
}

//SendHTMLEmail: 发送HTML格式邮件
func SendHTMLEmail(receiver string, path string, variables interface{}) error {
	var emailConfig EmailConfig
	DB.Raw("select * from email where id = 1 limit 1").Find(&emailConfig)
	e := email.NewEmail()
	e.Subject = "Micah在线视频后台管理"
	e.From = emailConfig.User
	e.To = strings.Split(receiver, ",")
	t := template.Must(template.ParseFiles(path))
	body := new(bytes.Buffer)
	//作为变量传递给html模板
	t.Execute(body, variables)
	e.HTML = body.Bytes()
	auth := smtp.PlainAuth("", emailConfig.User, emailConfig.Pass, emailConfig.Host)
	err := e.Send(emailConfig.Host+emailConfig.Port, auth)
	log.Println("failed to send email:", err)
	return err
}

//SendEmailCode: 发送邮箱验证码
func SendEmailCode(duration time.Duration, email, redisKey, templatePath string) *BusinessException {
	var code string
	if c, err := redis.RedisClient.Get(context.Background(), redisKey).Result(); err == nil {
		code = c
	} else {
		code = util.GetVerifyCode()
		redis.RedisClient.Set(context.Background(), redisKey, code, duration)
	}
	if err := SendHTMLEmail(email, templatePath, code); err != nil {
		return NewBusinessException(SEND_EMAIL_CODE_ERROR)
	}
	return nil
}
