package helpers

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func Hash(pass string) string {
	bytePass := []byte(pass)
	hash, err := bcrypt.GenerateFromPassword(bytePass, bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func VerifyPassword(hashedPass string, plainPass string) bool {
	byteHash := []byte(hashedPass)
	bytePlain := []byte(plainPass)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePlain)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
