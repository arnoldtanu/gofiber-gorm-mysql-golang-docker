package services

import (
	"errors"

	"github.com/arnoldtanu/disbursement-api/src/models"
	"github.com/arnoldtanu/disbursement-api/src/repositories"
	"golang.org/x/crypto/bcrypt"
)

func DoDisbursement(userid, amount int, passkey string) (*models.Users, error) {
	user, err := repositories.GetUserBalanceAndPasskey(userid);

	if err!=nil{
		return nil, errors.New(err.Error());
	}

	match := CheckPasswordHash(passkey, user.Passkey)
	if !match {
		return nil, errors.New("wrong passkey");
	}

	if user.Balance < amount {
		return nil, errors.New("insufficient balance");
	}

	user.Balance = user.Balance - amount;
	if errorUpdateUserBalance := repositories.UpdateUserBalance(&user); errorUpdateUserBalance!=nil{
		return nil, errors.New(errorUpdateUserBalance.Error())
	}

	return &user,nil;
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}