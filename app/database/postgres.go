package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabasePostgres() (*gorm.DB, error) {

	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	dsn := fmt.Sprintf("user=%s password=%s dbname=%s search_path=%s", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_SCHEMA"))
	db, err := gorm.Open(postgres.Open(dsn))

	return db, err

}
