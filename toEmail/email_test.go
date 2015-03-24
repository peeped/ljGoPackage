package toEmail

import (
	"testing"
)

var mailbody = `<html>
				<head><title>%s-订单审核</title></head>
				<body>
				<p>你猜中国字<p>
				</body>
				</html>`

func TestSendMail(t *testing.T) {
	mail, err := NewEmail("html", "sender@email.com", "iliaojie@gmail.com", "测试各邮件都没有乱码", mailbody)
	if err != nil {
		panic(err)
	}
	t.Logf(" %v\n", mail.SendMail("sender@email.com", "sender pass", "smtp.xxx.xxx.com", "25"))
}
