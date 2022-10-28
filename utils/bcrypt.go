package utils

import (
	"encoding/base64"
	"golang.org/x/crypto/bcrypt"
	"regexp"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func IsPhone(phone string) bool {
	v := regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)
	return v.MatchString(phone)
}

func EncodeString(str string) string {
	encode := base64.StdEncoding.EncodeToString([]byte(str))
	return encode
}

func DecodeString(str string) string {
	decode, _ := base64.StdEncoding.DecodeString(str)
	return string(decode)
}
