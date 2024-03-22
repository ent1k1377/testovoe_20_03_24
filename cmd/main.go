package main

import (
	"database/sql"
	db "github.com/ent1k1377/testovoe_20_03_24/db/sqlc"
	"github.com/ent1k1377/testovoe_20_03_24/internal/app"
	"github.com/ent1k1377/testovoe_20_03_24/internal/util"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func main() {
	// Загрузка переменных окружения из файла .env
	err := loadEnvironment()
	if err != nil {
		log.Fatal("cannot load environment", err)
	}

	// Создание нового экземпляра конфигурации
	config := util.NewConfig()

	// Подключение к базе данных
	conn, err := sql.Open(config.DatabaseDriver, config.DatabaseURL)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}
	defer conn.Close()

	// Создание хранилища для взаимодействия с базой данных
	store := db.NewStore(conn)

	// Запуск приложения
	if err := app.StartProcessingOrders(store, os.Args[1:]); err != nil {
		log.Fatal("app doesnt work", err)
	}
}

func loadEnvironment() error {
	return godotenv.Load()
}
