package functions

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gopkg.in/gomail.v2"
	"serverTest/conf"
	"time"
)

type SendEmailBody struct {
	SenderName       string
	UserEmail        string
	Password         string
	Receivers        []string
	IsDefaultSender  bool
	Title            string
	Body             string
	Host             string
	Port             int
}

func NewDefaultEmailBody(title, body string, receivers []string) *SendEmailBody{
	return &SendEmailBody{
		SenderName: "little fairy",
		Receivers: receivers,
		IsDefaultSender: true,
		Title: title,
		Body: body,
		Host: "smtp.qq.com",
		Port: 465,
	}
}

func DefaultEmailIsInvalid() bool {
	cfg := conf.Cfg
	flagA := cfg.Email == nil || cfg.Email.UserEmail == "" || cfg.Email.Password == ""
	flagB := cfg.Email.Sender == "" || cfg.Email.Host == "" || cfg.Email.Port == 0

	return flagA || flagB
}

func SendEmail(emailBody *SendEmailBody) error {
	m := gomail.NewMessage()

	if emailBody.IsDefaultSender {
		cfg := conf.Cfg
		if DefaultEmailIsInvalid() {
			log.Errorf("无法使用默认发送人，请检查配置")
			return fmt.Errorf("无法使用默认发送人，请检查配置")
		}

		emailBody.UserEmail = cfg.Email.UserEmail
		emailBody.Password = cfg.Email.Password
		emailBody.Host = cfg.Email.Host
		emailBody.Port = cfg.Email.Port
		emailBody.SenderName = cfg.Email.Sender
	}

	m.SetHeader("From", m.FormatAddress(emailBody.UserEmail, emailBody.SenderName))
	m.SetHeader("To", emailBody.Receivers...)

	m.SetHeader("Subject", emailBody.Title)
	m.SetBody("text/html", emailBody.Body)

	d := gomail.NewDialer(emailBody.Host, emailBody.Port, emailBody.UserEmail, emailBody.Password)

	if err := d.DialAndSend(m); err != nil {
		log.Errorf("发送邮件失败,err=%+v", err.Error())
		return err
	}

	return nil
}

func SetDelaySendEmail(emailBody *SendEmailBody, second int) {
	log.Infof("一封信的邮件在:%+vs后发送,邮件内容为:%+v", second, emailBody)

	time.Sleep(time.Second * time.Duration(second))
	err := SendEmail(emailBody)
	if err != nil {
		log.Errorf("延时邮件发送失败,请检查报错:%+v", err.Error())
		return
	}

	log.Infof("延时邮件发送成功")
}
