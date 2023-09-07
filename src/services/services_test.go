package services

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestCheckPasswordHash(t *testing.T){
	originalString := "HelloWorld"
	temp, _ := bcrypt.GenerateFromPassword([]byte(originalString), 14)
	hashedString := string(temp)

	result := CheckPasswordHash(originalString,hashedString)
	if !result {
		panic("hashing the same string twice, results non-matching value")
	}

	result = CheckPasswordHash("Hi"+originalString+"Dunia",hashedString)
	if (result){
		panic("hashing two different string, results matching value")
	}
}