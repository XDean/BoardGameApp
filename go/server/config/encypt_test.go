package config

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var key = "123456"

func TestGenerate(t *testing.T) {
	fmt.Println(EncryptString("data-here", "key-here"))
}

func TestDecrypt(t *testing.T) {
	encrypted := Encrypt([]byte("abc"), key)
	decrypted, err := Decrypt(encrypted, key)
	assert.NoError(t, err)
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
			AppId:     EncryptString("appid", key),
			AppSecret: EncryptString("appseceret", key),
		},
		Security: Security{
			Key: EncryptString("key", key),
		},
	}

	err := Decode(&input, key)
	assert.NoError(t, err)
	assert.Equal(t, expected, input)
}
