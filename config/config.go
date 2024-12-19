package config

import (
    "log"
    "os"

    "github.com/joho/godotenv"
)

func LoadEnv() {
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found, using default env variables")
    }
}

func GetDBConnStr() string {
    return "user=" + os.Getenv("DB_USER") +
        " password=" + os.Getenv("DB_PASSWORD") +
        " dbname=" + os.Getenv("DB_NAME") +
        " host=" + os.Getenv("DB_HOST") +
        " port=" + os.Getenv("DB_PORT") +
        " sslmode=disable"
}
