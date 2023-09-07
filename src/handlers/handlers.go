package handlers

import (
	"strconv"

	"github.com/arnoldtanu/disbursement-api/src/services"
	"github.com/gofiber/fiber/v2"
)

type DisbursementParams struct{
	ID      int    `json:"id"`
	Amount  int    `json:"amount"`
	Passkey string `json:"passkey"`
}

func HelloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func Disbursement(c *fiber.Ctx) error {
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

	user, err := services.DoDisbursement(disbursementParams.ID, disbursementParams.Amount, disbursementParams.Passkey);
	if err!=nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "failed",
			"message": err.Error(),
		});
	} else {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status": "success",
			"message": "successful disbursement, updated balance : "+ strconv.Itoa(user.Balance),
		});
	}
}