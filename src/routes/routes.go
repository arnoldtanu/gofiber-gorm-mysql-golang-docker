package routes

import (
	"github.com/arnoldtanu/disbursement-api/src/handlers"
	"github.com/gofiber/fiber/v2"
)

func TransactionRoutes(app *fiber.App){
  app.Get("/",handlers.HelloWorld)
  app.Post("/disbursement",handlers.Disbursement)
}