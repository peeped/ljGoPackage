package common

import (
	"encoding/base64"
	"fmt"
	"github.com/peeped/ljGoPackage/common"
	"testing"
)

func TestAesEncrypt(t *testing.T) {
	// AES-128。key长度：16, 24, 32 bytes 对应 AES-128, AES-192, AES-256
	aesKey := []byte("d7b59483d9954929b3262acc677cb51a")

	// AES加密
	result, err := common.AesEncrypt([]byte("abcd@163.com"), aesKey)
	if err != nil {
		panic(err)
	}
	t.Logf("加密后并BASE64:%v\n", base64.StdEncoding.EncodeToString(result))

	// AES解密
	origData, err := common.AesDecrypt(result, aesKey)
	if err != nil {
		panic(err)
	}
	t.Logf("解密后：:%v\n", string(origData))
}
