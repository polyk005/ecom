package main

import (
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/v4/database/mysql"
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

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatalf(err)
	}

	m, err := migrate.NewWithSourceInstance(
		"file://cmd/migrate/migrations",
		"postgresql",
		driver,
	)

	if err != nil {
		log.Fatalf(err)
	}

	cmd := os.Args[(len(os.Args) - 1)]
	if cmd == "up" {
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}

	if cmd == "down" {
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf(err)
		}
	}
}
