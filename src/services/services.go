package services

import (
	"github.com/arnoldtanu/disbursement-api/src/repositories"
	"golang.org/x/crypto/bcrypt"
)

func DoDisbursement(userid, amount int, passkey string) (string, error) {
	user, errorGetUser := repositories.GetUserBalanceAndPasskey(userid);

	if errorGetUser!=nil{
		return "user not found", nil;
	}

	match := CheckPasswordHash(passkey, user.Passkey)
	if !match {
		return "wrong passkey",nil;
	}

	if user.Balance < amount {
		return "insufficient balance",nil;
	}

	user.Balance = user.Balance - amount;
	if errorUpdateUserBalance := repositories.UpdateUserBalance(&user); errorUpdateUserBalance!=nil{
		return errorUpdateUserBalance.Error(), nil
	}

	return "",nil;
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}