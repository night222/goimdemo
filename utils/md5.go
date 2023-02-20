package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

// MD5加密小写
func Md5Encode(data string) string {
	h := md5.New()
	//加密
	h.Write([]byte(data))
	//获取加密切片
	TempData := h.Sum(nil)
	return hex.EncodeToString(TempData)
}

// MD5加密大写
func MD5Encode(plainpwd, salt string) string {
	return strings.ToUpper(Md5Encode(plainpwd + salt))
}

// 密码加密
func MakePassword(plainpwd, salt string) string {
	return MD5Encode(plainpwd, salt)
}

// 密码验证
func ValidPassword(plainpwd, password, salt string) bool {
	if MakePassword(plainpwd, salt) == password {
		return true
	}
	return false
}
