package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"shop/database/models"
)

func New() (*gorm.DB, error) {
	log.Println("connecting...")
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable search_path=shop",
		"localhost", 5432, "postgres", "changeme", "shop",
	)

	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return nil, err
	}
	log.Println("successfully connected to db!")
	if err = db.AutoMigrate(
		models.Comment{},
		models.Product{},
		models.Credentials{},
		models.User{},
	); err != nil {
		return nil, err
	}

	return db, nil
}
