package common

import (
	"encoding/base64"
	"github.com/peeped/ljGoPackage/common"
	"testing"
)

var privateKey = []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAu9n07e4eE7Biacx9R0LEiQ3rZgcwz9LJjyKol7EdhlVHJ98J
qhYXfotJdeXGLQvxGADmsbDY5ldr0Wk1UsjNwxgTCXIXSabXKc7rOBlX1PaiPiAF
3qjc5rv660s0K9oWckZ+08i4NNmQFsZ6a0HgsyfSLG2hfegWh77Orgcz2BcnRoBL
9MJCEuFENqf+xdueVTfZz2rCM7dugMcZTyoWFcP74c76yw4vHDQAPjpbmrh0WuMI
M5cJ4W+u/856kEaD60mxyPEL+1W2NJQaJ9usSNDYvXKr3Vonu0I0BQt0LfuYHzWK
3iBwq3SzcamR/60heziJgZ0bJBofQyHj+Cn9lwIDAQABAoIBACqS4izOa6igsB00
SCxPWIWLTw9nj8t6BU5YV4dRj9RzHVZO+TzAFwEKBlMfCUQKUdDT23ToFLBXncrv
IjOp6OBPY3kfj2GU22zRRYQIUlykrO6RiWMGOFJexiZve9p4ad/qVDIhaoYnzL0s
rHAElS1lV//TtOb5I6oON38/iKNvcEKJlhxELvqzLDOnGnAKzM3l93J5dmrkpmOZ
8rZVyeXNtrHisrrJ0AoDmj7MXr4mmmiNs7ZJb2yv55B+fq0JyR/icuI0jfcgtd0V
Boh2gSULZObfgD0j9IMYlwkiR4DmkeOuldewNcz4l44tCy9m+FS7nBa6kYo1J/7P
4zWg8zECgYEAzCmi8ODfdKZ7q4am/1zqI9dEfA6dGEIQQyREGlS65blzOlFLq+hG
2QMtzL4MenRUpP/QRLzUnkiJqupds2iWrrfte6Kec3uo4rItrSnnaXcLWH7dNOeD
WBSeeKcIfkBHpPRDkuiC8okbhf8Ng4sqRT5+oUM73N/7uOjRkAthNV8CgYEA64wb
ELEEgtAIeWzZTXekILv8RWREYmSYjHrBG4zKSsouhIXjl7t0eHxXfzWtLBBkLfWJ
X8pvhawt74g2EhYMAYBx7q4CwfBDcrNisTgfoep19a7d3pHM7IUHSEXXhBoIFJDi
q+24bGmyZaPBUEUlJ7MeTrh3ILX/2tk9E0haqskCgYBKEITC++E0sTzGIggtNajf
LbXzh12oMjcyFFL8dmaC9j7+FgXsrEwfaA7SatOeDNu0K/WDKjm73jbLIVCyyCt5
4NGve3QeEutWqir12fDQitY72XIoQiCc8IX44SesnWcgSVjGT8FJeUHZ34gog3Dn
Q9+uYvSxkTQBhbyYk/hE4wKBgHJeQ+H14XfWpNa4aEZ5+gI+5H2Y8q9Hot5K2CqV
UL/BrZaBIAHTbfj2ftFwcZX8m3fJSZtuQnoIIQG2BHMBq3CrOiam7QXXsBgoS5o6
4vkOS5ov/uCLsJGDAgcwijVFInlB5B2QvkQ9ifZZ7YoZGLJPAT89x/HlDMbpRgNv
1T4pAoGBAIvPw0EyjjmzyQrGGoZTcrpS5u4qmM0dKGkBF21QFnQlnrgRNPeMgwLe
h78Pfag1//hY3qm3hE9t23CUlyxxrsBmri6yYUMT2kspIhoAGCeIHBqF35dBmj8Q
ypHxVH9XYbZRUgR7LWtpXOD+9ySUG3VoK0OJTpUN+sLDfVl5T59g
-----END RSA PRIVATE KEY-----

`)

var publicKey = []byte(`-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAu9n07e4eE7Biacx9R0LE
iQ3rZgcwz9LJjyKol7EdhlVHJ98JqhYXfotJdeXGLQvxGADmsbDY5ldr0Wk1UsjN
wxgTCXIXSabXKc7rOBlX1PaiPiAF3qjc5rv660s0K9oWckZ+08i4NNmQFsZ6a0Hg
syfSLG2hfegWh77Orgcz2BcnRoBL9MJCEuFENqf+xdueVTfZz2rCM7dugMcZTyoW
FcP74c76yw4vHDQAPjpbmrh0WuMIM5cJ4W+u/856kEaD60mxyPEL+1W2NJQaJ9us
SNDYvXKr3Vonu0I0BQt0LfuYHzWK3iBwq3SzcamR/60heziJgZ0bJBofQyHj+Cn9
lwIDAQAB
-----END PUBLIC KEY-----`)

func Test(t *testing.T) {

	// rsa加密
	data, err := common.RsaEncrypt([]byte("email=cbljc@163.com&name=沙子"), publicKey)
	if err != nil {
		panic(err)
	}
	t.Log("rsa加密结果", base64.StdEncoding.EncodeToString(data))

	// rsa解密

	origData, err := common.RsaDecrypt(data, privateKey)
	if err != nil {
		panic(err)
	}
	t.Log("rsa解密结果", string(origData))

}
