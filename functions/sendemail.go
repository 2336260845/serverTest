package functions

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gopkg.in/gomail.v2"
	"serverTest/conf"
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
