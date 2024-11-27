package main

import (
	"fmt"
	"log"

	"github.com/sikozonpc/ecom/cmd/api"
	"github.com/sikozonpc/ecom/config"
	"github.com/sikozonpc/ecom/db"
	"gorm.io/gorm"
)

func main() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require",
		config.Envs.DBAddress,
		config.Envs.DBUser,
		config.Envs.DBPassword,
		config.Envs.DBName,
		"5432",
	)

	db, err := db.NewPostgresStorage(dsn)
	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)

	server := api.NewAPIServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *gorm.DB) {
	err := db.Exec("SELECT 1")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("DB: Successfully connected!")
}
