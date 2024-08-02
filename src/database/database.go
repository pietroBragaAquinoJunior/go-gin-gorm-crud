package database

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "log"
    "github.com/pietroBragaAquinoJunior/go-gin-gorm-crud/src/models"
)

var DB *gorm.DB

func Connect() {
    dsn := "user:user_password@tcp(127.0.0.1:3306)/mydatabase?charset=utf8mb4&parseTime=True&loc=Local"
    var err error
    DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Error connecting to database: %v", err)
    }
    log.Println("Database connection established")

    // Auto migrate
    err = DB.AutoMigrate(&models.Product{})
    if err != nil {
        log.Fatalf("Error migrating database: %v", err)
    }
}
