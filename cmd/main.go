package main

import (
	"database/sql"
	"github.com/ent1k1377/testovoe_20_03_24/internal"
	"github.com/joho/godotenv"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	err := loadEnvironment()
	if err != nil {
		log.Fatal("cannot load environment", err)
	}

	config := internal.NewConfig()

	conn, err := sql.Open(config.DatabaseDriver, config.DatabaseURL)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}
	if conn.Ping() == nil {
		log.Fatal("connection not established")
	}
	log.Printf("All good!")
}

func loadEnvironment() error {
	return godotenv.Load()
}
