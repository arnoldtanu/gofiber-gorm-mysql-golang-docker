package models

import (
	"gorm.io/gorm"
)

type Users struct{
	gorm.Model
	id       int    `gorm:"column:id;"`
	Username string `gorm:"column:username;"`
	Passkey  string `gorm:"column:passkey;"`
	Balance  int    `gorm:"column:balance;"`
}