package main

import (
	"fmt"
	"os"

	"github.com/arnoldtanu/disbursement-api/src/routes"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
func main() {
    password := "secret"
    hash, _ := HashPassword(password) // ignore error for the sake of simplicity

    fmt.Println("Password:", password)
    fmt.Println("Hash:    ", hash)

    // dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
    dsn := fmt.Sprintf(
      "%s:%s@tcp(db-disbursement:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
      os.Getenv("MYSQL_USER"),
      os.Getenv("MYSQL_PASSWORD"),
      os.Getenv("MYSQL_DATABASE"),
    )
    _, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    fmt.Println("dsn:     ",dsn)
    fmt.Println("Error?:  ",err)

  app := fiber.New()

  routes.TransactionRoutes(app)

  app.Listen(":3000")

}

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}