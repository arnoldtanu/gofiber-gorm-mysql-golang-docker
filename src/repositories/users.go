package repositories

import (
	"errors"

	db "github.com/arnoldtanu/disbursement-api/src/config"
	"github.com/arnoldtanu/disbursement-api/src/models"
)

func GetUserBalanceAndPasskey(userid int) (models.Users, error) {
	var user models.Users

	if err:=db.DB.Db.First(&user,userid).Error; err!=nil{
		return user, errors.New("user not found")
	}

	return user, nil;
}

func UpdateUserBalance(user *models.Users) error {
	if err:=db.DB.Db.Model(&user).Update("Balance",user.Balance).Error; err!=nil{
		return errors.New("cannot update user's balance")
	}

	return nil;
}