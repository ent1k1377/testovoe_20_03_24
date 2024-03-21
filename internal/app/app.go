package app

import (
	"context"
	"fmt"
	"sort"
	"strings"

	db "github.com/ent1k1377/testovoe_20_03_24/db/sqlc"
)

// order представляет информацию о заказе.
type order struct {
	ProductID   int64
	ProductName string
	OrderID     int64
	Quantity    int32
	ShelveNames []string
}

// StartProcessingOrders начинает обработку заказов.
func StartProcessingOrders(sqlStore *db.SQLStore, orderNumbers []int64) error {
	ctx := context.Background()

	// Получаем информацию о заказах из хранилища данных SQL.
	ordersInfo, err := sqlStore.Queries.GetOrderInfo(ctx, orderNumbers)
	if err != nil {
		return fmt.Errorf("ошибка при получении информации о заказах: %w", err)
	}

	return processOrders(ordersInfo)
}

// processOrders обрабатывает информацию о заказах.
func processOrders(ordersInfo []db.GetOrderInfoRow) error {
	// Создаем карту для хранения дополнительных стеллажей для каждого продукта.
	additionalShelves := make(map[int64][]string)

	// Создаем карту для хранения заказов по стеллажам.
	orders := make(map[string][]order)

	// Обработка информации о заказах.
	for _, o := range ordersInfo {
		if !o.ShelveIsPrimary {
			additionalShelves[o.ProductID] = append(additionalShelves[o.ProductID], o.ShelveName)
			continue
		}

		additionalShelves[o.ProductID] = []string{o.ShelveName}

		orders[o.ShelveName] = append(orders[o.ShelveName], order{
			ProductID:   o.ProductID,
			ProductName: o.ProductName,
			OrderID:     o.OrderID,
			Quantity:    o.Quantity,
			ShelveNames: make([]string, 0),
		})
	}

	processAdditionalShelves(orders, additionalShelves)
	return nil
}

// processAdditionalShelves обрабатывает дополнительные стеллажи для каждого продукта.
func processAdditionalShelves(orders map[string][]order, additionalShelves map[int64][]string) {
	for _, v := range orders {
		for i := range v {
			v[i].ShelveNames = append(v[i].ShelveNames, additionalShelves[v[i].ProductID][1:]...)
		}
	}

	printOrders(orders)
}

// printOrders выводит информацию о заказах.
func printOrders(orders map[string][]order) {
	var orderKeys []string

	for k := range orders {
		orderKeys = append(orderKeys, k)
	}
	sort.Strings(orderKeys)

	for _, e := range orderKeys {
		fmt.Printf("===Стеллаж %s\n", e)
		v := orders[e]
		for _, o := range v {
			fmt.Printf("%s (id=%d)\nзаказ %d, %d шт\n",
				o.ProductName,
				o.ProductID,
				o.OrderID,
				o.Quantity)
			if len(o.ShelveNames) != 0 {
				fmt.Printf("доп стеллаж: %s\n", strings.Join(o.ShelveNames, ","))
			}
			fmt.Println()
		}
	}
}
