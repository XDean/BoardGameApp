package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var key = "123456"

func TestDecrypt(t *testing.T) {
	encrypted := Encrypt([]byte("abc"), key)
	decrypted := Decrypt(encrypted, key)
	assert.Equal(t, "abc", string(decrypted))
}

func TestDecode(t *testing.T) {
	expected := Config{
		Debug: true,
		Wechat: Wechat{
			AppId:     "appid",
			AppSecret: "appseceret",
		},
		Security: Security{
			Key: "key",
		},
	}
	input := Config{
		Debug: true,
		Wechat: Wechat{
			AppId:     encryptString("appid"),
			AppSecret: encryptString("appseceret"),
		},
		Security: Security{
			Key: encryptString("key"),
		},
	}

	err := Decode(&input, key)
	assert.NoError(t, err)
	assert.Equal(t, expected, input)
}

func encryptString(s string) string {
	return "ENC~" + string(Encrypt([]byte(s), key))
}
