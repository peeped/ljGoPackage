package email

import (
	"testing"
)

func TestSendMail(t *testing.T) {
	mailbody := `<html>
<head>
<title>%s-订单审核</title>

</head>
<body>
    <p>你猜中国字<p>
</body>
</html>`
	mail := NewMail("html", "support@zcgames.cn", "liaojie@zcgames.cn", "再次测试各邮件是否有乱码", mailbody)
	// email := New("html", "support@zcgames.cn", "liaojie@zcgames.cn", "再次测试各邮件是否有乱码", mailbody)
	// t.Logf(" %v\n", email.SendMail("support@zcgames.cn", "zcxy1234", "smtp.exmail.qq.com", "25"))
}
