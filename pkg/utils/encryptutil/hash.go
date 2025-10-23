package encryptutil

import (
	"crypto/md5"
	"encoding/hex"

	"golang.org/x/crypto/bcrypt"
)

func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func HashPassword(value string) string {
	hashedBytes, _ := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)
	return string(hashedBytes)
}

// VerifyPassword 验证加密的文本是否与纯文本相同
func VerifyPassword(hash, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
