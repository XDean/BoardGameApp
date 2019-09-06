package config

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"io"
	"reflect"
	"strings"
)

func Decode(obj interface{}, key string) error {
	var v reflect.Value
	if ov, ok := obj.(reflect.Value); ok {
		v = ov
	} else {
		v = reflect.ValueOf(obj)
	}
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		for field.Kind() == reflect.Ptr {
			field = field.Elem()
		}
		switch field.Kind() {
		case reflect.String:
			str := field.String()
			if strings.HasPrefix(str, "ENC~") {
				text, err := Decrypt(str[4:], key)
				if err != nil {
					return err
				}
				field.SetString(string(text))
			}
		case reflect.Struct:
			err := Decode(field, key)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func Encrypt(data []byte, passphrase string) string {
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
	seal := gcm.Seal(nonce, nonce, data, nil)
	return base64.StdEncoding.EncodeToString(seal)
}

func Decrypt(text string, passphrase string) ([]byte, error) {
	data, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		return nil, errors.New("Invalid text to decrypt")
	}
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
	return plaintext, nil
}

func EncryptString(s string, key string) string {
	return "ENC~" + string(Encrypt([]byte(s), key))
}

func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}
