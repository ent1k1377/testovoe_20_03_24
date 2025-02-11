package main

import (
	"database/sql"
	"github.com/ent1k1377/testovoe_20_03_24/internal/app"
	"github.com/ent1k1377/testovoe_20_03_24/internal/repository/postgres"
	"github.com/ent1k1377/testovoe_20_03_24/internal/util"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
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
	repo := postgres.NewRepository(conn)

	// Создание приложения
	a := app.NewApp(repo)

	// Запуск приложения
	if err := a.Run(); err != nil {
		log.Fatal("app doesnt work: ", err)
	}
}

func loadEnvironment() error {
	return godotenv.Load()
}
