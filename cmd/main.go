package main

import (
	"database/sql"
	"github.com/ent1k1377/testovoe_20_03_24/internal"
	"github.com/joho/godotenv"
	"log"
	"strconv"
	"time"

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

	var currentDate time.Time
	err = conn.QueryRow("select current_date").Scan(&currentDate)
	log.Printf("current date: %s\n", currentDate.Format("2006-01-02"))

	log.Printf("All good!")
}

func recordData(arg []string) []int {
	newA := make([]int, len(arg))
	for i, _ := range arg {
		elem, err := strconv.Atoi(arg[i])
		if err == nil {
			return nil
		}
		newA[i] = elem
	}
	return nil
}

func loadEnvironment() error {
	return godotenv.Load()
}
