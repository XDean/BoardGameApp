package config

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"io"
	"io/ioutil"
	"reflect"
	"strings"

	"gopkg.in/yaml.v2"
)

// Global instance
var Global Config

// Conf is the root configuration struct
type Config struct {
	Debug    bool
	Security Security
	DB       DB
	Wechat   Wechat
}

type Security struct {
	Key string
}

// DB is Database configuration struct
type DB struct {
	Dialect string
	URL     string
}

type Wechat struct {
	AppId     string
	AppSecret string
	AuthUrl   string
}

// Init configuration instance
func (c *Config) Load(path string) (err error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}
	err = yaml.Unmarshal(content, c)
	if err == nil {
		err = c.decode()
	}
	return
}

func Decode(obj interface{}, key string) error {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		for field.Kind() == reflect.Ptr {
			field = field.Elem()
		}
		switch field.Kind() {
		case reflect.String:
			str := field.String()
			if strings.HasPrefix(str, "ENC~") {
				field.SetString(string(Decrypt([]byte(str), key)))
			}
		case reflect.Struct:
			field.set
		}
	}
}

func Encrypt(data []byte, passphrase string) []byte {
	block, err := aes.NewCipher([]byte(createHash(passphrase)))
	if err != nil {
		panic(err.Error())
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	return gcm.Seal(nonce, nonce, data, nil)
}

func Decrypt(data []byte, passphrase string) []byte {
	key := []byte(createHash(passphrase))
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}
	return plaintext
}

func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}
