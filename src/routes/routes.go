package routes

import (
	"github.com/arnoldtanu/disbursement-api/src/handlers"
	"github.com/gofiber/fiber/v2"
)

func TransactionRoutes(app *fiber.App){
  app.Get("/",handlers.HelloWorld)
}

/*
endpoint: /disbursment
method: post
params:
  - userID : int
    amount : int
    passkey: string //kalau ada waktu
output:
  json {
    status: success/error,
    data: {
      currBalance: int
    }
  }
validation:
  1.check value amount sudah benar
  2.check passkey (password kedua untuk withdrawal), sudah sama kayak yg disimpan di database nggak (kl nutut)
  3.check balance user tersebut lbh besar sama dengan amount
logic: 
  kalau validation sukses semua, potong balance user tersebut sesuai amount, lalu return sukses
*/