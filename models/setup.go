package models

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB


func ConnDatabase() {


  err := godotenv.Load()
  if err != nil {
		log.Fatal("Error loading .env file")
	}

  dsn := os.Getenv("DB_CONNECTION_STRING")
	if dsn == "" {
		log.Fatal("DB_CONNECTION_STRING not set in .env file")
	}


  db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

  if err != nil {
    panic("failed to connect db...")
  }

  db.AutoMigrate(&Character{})

  DB = db



} 
