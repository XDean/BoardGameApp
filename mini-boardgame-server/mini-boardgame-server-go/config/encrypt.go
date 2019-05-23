package config

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
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
				field.SetString(string(Decrypt([]byte(str[4:]), key)))
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
