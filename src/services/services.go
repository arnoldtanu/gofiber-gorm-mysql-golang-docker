package services

import (
	db "github.com/arnoldtanu/disbursement-api/src/config"
	"github.com/arnoldtanu/disbursement-api/src/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type DisbursementParams struct{
	ID      int    `json:"id"`
	Amount  int    `json:"amount"`
	Passkey string `json:"passkey"`
}

func DoDisbursement(c *fiber.Ctx) error {
	disbursementParams := new(DisbursementParams)
	if err := c.BodyParser(disbursementParams); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "failed",
			"message": err.Error(),
		})
	}

	var missingRequiredData string
	if disbursementParams.Amount <= 0 { missingRequiredData = "Amount"
	} else if disbursementParams.ID == 0 { missingRequiredData = "id"
	} else if disbursementParams.Passkey == "" { missingRequiredData = "passkey"	}

	if len(missingRequiredData) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "failed",
			"message": missingRequiredData+" is required or wrong value",
		});
	}

	var user models.Users
	if err:=db.DB.Db.First(&user,disbursementParams.ID).Error; err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "failed",
			"message": "user not found",
		});
	}

	match := CheckPasswordHash(disbursementParams.Passkey, user.Passkey)
	if !match {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "failed",
			"message": "wrong passkey",
		});
	}

	if user.Balance < disbursementParams.Amount {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "failed",
			"message": "insufficient balance",
		});
	}

	user.Balance = user.Balance - disbursementParams.Amount;
	db.DB.Db.Model(&user).Update("Balance",user.Balance)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"currBalance": user.Balance,
	});
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}