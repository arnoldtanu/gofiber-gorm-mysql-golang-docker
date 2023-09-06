package handlers

import (
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
	return services.DoDisbursement(c)
}