package config

import (
	"fmt"
	"user-rest-app/models"

	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDB() *gorm.DB {
	// connect to db
	// db related crede  ntials, connection string, dbname etc
	// .env file
	// read details from .env file
	err:=godotenv.Load(".env")
	if err!=nil{
		log.Fatalf("Error loading .env file")
		panic("Error loading .env file")
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
        os.Getenv("DB_HOST"),
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_NAME"),
        os.Getenv("DB_PORT"),
    )
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database")
		//panic("Error connecting to database")
	}

	//db.AutoMigrate(&models.Employee{})
	return db
	}

func MigrateDB(db *gorm.DB) error {
	return db.AutoMigrate(&models.Role{}, &models.User{})
}