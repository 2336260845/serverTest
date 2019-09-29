package functions

import "testing"

func TestSendEmail(t *testing.T) {
	emailBody := SendEmailBody{
		SenderName: "little fairy",
		UserEmail: "853266983@qq.com",
		Password: "woheljvcosjtbaif",
		Receivers: []string{"2336260845@qq.com"},
		IsDefaultSender: false,
		Title: "你有事情需要处理",
		Body: "这是一封测试邮件",
		Host: "smtp.qq.com",
		Port: 465,
	}

	err := SendEmail(&emailBody)
	if err != nil {
		t.Errorf("发送邮件失败,err=%+v", err.Error())
		return
	}

	t.Log("测试通过")
}
