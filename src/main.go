package main

import (
	db "github.com/arnoldtanu/disbursement-api/src/config"
	"github.com/arnoldtanu/disbursement-api/src/routes"
	"github.com/gofiber/fiber/v2"
)
func main() {
  db.ConnectDb()

  app := fiber.New()
  routes.TransactionRoutes(app)
  app.Listen(":3000")
}