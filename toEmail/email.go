package email

import (
	"encoding/base64"
	"fmt"
	"net/smtp"
	"strings"
)

type Mail struct {
	ContentType string
	From        string
	To          string
	Subject     string
	Body        []byte
}

// Format address for use in headers.
func format(mail string) string {
	return fmt.Sprintf("=?utf-8?b??= <%v>", mail)
}

func encode(x string) string {
	return base64.StdEncoding.EncodeToString([]byte(x))
}

func hencode(x string) string {
	return fmt.Sprintf("=?utf-8?b?%v?=", encode(x))
}

func NewEmail(contentType, from, to, subject, msg string) *Mail {
	return &Mail{contentType, from, to, subject, []byte(msg)}
}

func (m Mail) SendMail(smtp_account, smtp_pass, smtp_host, port string) error {
	// msg := m.Message()
	auth := smtp.PlainAuth(m.From, smtp_account, smtp_pass, smtp_host)
	return smtp.SendMail(fmt.Sprintf("%s:%s", smtp_host, port), auth, smtp_account, strings.Split(m.To, ";"), m.message())
}
func (m Mail) message() (msg []byte) {
	content_Type := "html"
	if m.ContentType != content_Type {
		content_Type = "plain"
	}
	// msg = []byte(fmt.Sprintf(`Content-Type: text/%s; charset="ISO-8859-1" MIME-Version: 1.0 Content-Transfer-Encoding: base64`, content_Type))
	msg = []byte(fmt.Sprintf(`Content-Type: text/%s; charset="UTF-8,ISO-8859-1" MIME-Version: 1.0 Content-Transfer-Encoding: base64`, content_Type))
	msg = append(msg, []byte(fmt.Sprintf("\r\nFrom: %v\n", format(m.From)))...)

	var tos []string
	for _, addr := range strings.Split(m.To, ";") {
		tos = append(tos, format(addr))
	}
	// msg = append(msg, []byte("Accept-Charset=\"ISO-8859-1,UTF-8\"\r\n")...)
	msg = append(msg, []byte(fmt.Sprintf("To: %v\n", strings.Join(tos, ", ")))...)

	msg = append(msg, []byte(fmt.Sprintf("Subject: %v\n\n%v\n", hencode(m.Subject), fmt.Sprintf("%s", string(m.Body))))...)
	return
}
