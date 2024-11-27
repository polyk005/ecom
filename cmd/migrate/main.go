package main

import (
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/sikozonpc/ecom/config"
	"github.com/sikozonpc/ecom/db"
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
	defer db.Close(dsn) // Закрытие соединения с БД

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("Error creating driver: %v", err)
	}

	m, err := migrate.NewWithSourceInstance(
		"file://cmd/migrate/migrations",
		"postgresql",
		driver,
	)

	if err != nil {
		log.Fatalf("Error creating migration instance: %v", err)
	}

	if len(os.Args) < 2 {
		log.Fatal("Command 'up' or 'down' is required")
	}

	cmd := os.Args[1]
	if cmd == "up" {
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
		log.Println("Migrations applied successfully.")
	} else if cmd == "down" {
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
		log.Println("Migrations rolled back successfully.")
	} else {
		log.Fatalf("Unknown command: %s. Use 'up' or 'down'.", cmd)
	}
}
