package app

import (
	"github.com/ent1k1377/testovoe_20_03_24/internal/app/service"
	"github.com/ent1k1377/testovoe_20_03_24/internal/repository"
	"github.com/ent1k1377/testovoe_20_03_24/internal/util"
	"log"
	"os"
)

type App struct {
	storage repository.Storage
}

func NewApp(storage repository.Storage) App {
	return App{storage: storage}
}

func (a *App) Run() error {
	ordersId, err := util.ConvertStringsToIntegers(os.Args[1:])
	if err != nil {
		log.Fatal("Incorrect transmitted data", err)
	}

	return service.HandleOrdersProcessing(a.storage, ordersId)
}
