package database

import (
    "fmt"
    "log"
    "os"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "github.com/joho/godotenv"
    "github.com/PMaini532/Sprint_Planner/models"
)

var DB *gorm.DB

func Connect() {

    err := godotenv.Load()
        if err != nil {
            log.Fatal("Error loading .env file")
        }

    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
        os.Getenv("DB_HOST"),
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_NAME"),
        os.Getenv("DB_PORT"),
    )


    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to DB:", err)
    }

    err = db.AutoMigrate(&models.User{})
    if err != nil {
        log.Fatal("Failed to run migrations:", err)
    }

    DB = db
    fmt.Println("🟢 Connected to PostgreSQL with GORM")
}
