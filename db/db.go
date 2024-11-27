package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresStorage(cfg string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(cfg), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}
