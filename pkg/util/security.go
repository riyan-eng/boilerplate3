package util

import "golang.org/x/crypto/bcrypt"

func GenerateHash(str string) string {
	hashedStr, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	PanicIfNeeded(err)
	return string(hashedStr)
}

func VerifyHash(hashedStr, candidateStr string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedStr), []byte(candidateStr)); err == nil {
		return true
	} else {
		return false
	}
}
