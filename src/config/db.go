package db

import (
	"fmt"
	"log"
	"os"

	"github.com/arnoldtanu/disbursement-api/src/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)
type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func ConnectDb(){
	dsn := fmt.Sprintf(
		"%s:%s@tcp(db-disbursement:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_DATABASE"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err!=nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	log.Println("DB connected!")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("running migration")
	db.AutoMigrate(&models.Users{})

	DB = Dbinstance{
		Db:db,
	}
}