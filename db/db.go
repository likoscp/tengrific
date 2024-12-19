package db

import (
	"log"
	"tengrific/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	dsn := config.GetDBConnStr()
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	log.Println("Database connected successfully")
}

func GetDBConnStr() string {
    return "user=postgres password=12345678 dbname=tengrific sslmode=disable"
}