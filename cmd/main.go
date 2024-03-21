package main

import (
	"database/sql"
	db "github.com/ent1k1377/testovoe_20_03_24/db/sqlc"
	"github.com/ent1k1377/testovoe_20_03_24/internal"
	"github.com/ent1k1377/testovoe_20_03_24/internal/util"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func main() {
	err := loadEnvironment()
	if err != nil {
		log.Fatal("cannot load environment", err)
	}

	config := util.NewConfig()

	conn, err := sql.Open(config.DatabaseDriver, config.DatabaseURL)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}

	store := db.NewStore(conn)

	ordersId, err := util.ConvertStringsToIntegers(os.Args[1:])
	if err != nil {
		log.Fatal("Incorrect transmitted data", err)
	}

	internal.Start(store, ordersId)
}

func loadEnvironment() error {
	return godotenv.Load()
}
