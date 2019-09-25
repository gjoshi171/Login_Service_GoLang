package service

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash), err
}

func CheckPasswordHash(hash string, password string) bool {
	byteHash := []byte(hash)
	bytePass := []byte(password)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePass)
	//fmt.Println("bhai error throw kiya ", err)
	return err == nil
}
