package service

import (
	"fmt"
	"github.com/ent1k1377/testovoe_20_03_24/internal/model"
	"github.com/ent1k1377/testovoe_20_03_24/internal/repository"
	"github.com/ent1k1377/testovoe_20_03_24/internal/repository/postgres"
	"os"
	"sort"
	"strings"
)

// HandleOrdersProcessing начинает обработку заказов.
func HandleOrdersProcessing(storage repository.Storage, ordersIds []int64) error {
	var ordersInfo []model.OrderInfo

	// Хранит key: id продукта и value: структуру Product
	productByID := make(map[int64]model.Product)

	// Хранит key: id продукта, value: количество продукта в заказе
	productOrderQuantity := make(map[int64]int32)

	shelves, err := storage.Shelf().GetAllShelves()
	if err != nil {
		return err
	}

	shelfNameByShelfID := make(map[int64]string)
	for _, v := range shelves {
		shelfNameByShelfID[v.ID] = v.Name
	}

	for _, id := range ordersIds {
		orderItems, err := storage.OrderItem().GetOrderItems(id)
		if err != nil {
			return err
		}

		products, err := storage.Product().GetProducts(postgres.GetOrderItemsIds(orderItems))
		if err != nil {
			return err
		}

		productsShelves, err := storage.ProductShelf().GetProductsShelves(postgres.GetProductsShelvesIds(products))
		if err != nil {
			return err
		}

		for _, v := range orderItems {
			productOrderQuantity[v.ProductID] = v.Quantity
		}
		for _, v := range products {
			productByID[v.ID] = model.Product{
				Name:     v.Name,
				Quantity: productOrderQuantity[v.ID],
			}
		}

		for _, v := range productsShelves {
			ordersInfo = append(ordersInfo, model.OrderInfo{
				ProductName:     productByID[v.ProductID].Name,
				ProductID:       v.ProductID,
				OrderID:         id,
				Quantity:        productByID[v.ProductID].Quantity,
				ShelveName:      shelfNameByShelfID[v.ShelvesID],
				ShelveIsPrimary: v.IsPrimary,
			})
		}
	}

	fmt.Printf("=+=+=+=\nСтраница сборки заказов %s\n\n", strings.Join(os.Args[1:], ","))
	return processOrders(ordersInfo)
}

// processOrders обрабатывает информацию о заказах.
func processOrders(ordersInfo []model.OrderInfo) error {
	// Создаем карту для хранения дополнительных стеллажей для каждого продукта.
	additionalShelves := make(map[int64][]string)

	// Создаем карту для хранения заказов по стеллажам.
	orders := make(map[string][]model.Order)

	// Обработка информации о заказах.
	for _, o := range ordersInfo {
		if !o.ShelveIsPrimary {
			additionalShelves[o.ProductID] = append(additionalShelves[o.ProductID], o.ShelveName)
			continue
		}

		additionalShelves[o.ProductID] = []string{o.ShelveName}

		orders[o.ShelveName] = append(orders[o.ShelveName], model.Order{
			OrderID:     o.OrderID,
			ProductID:   o.ProductID,
			ProductName: o.ProductName,
			Quantity:    o.Quantity,
			ShelveNames: make([]string, 0),
		})
	}

	groupAdditionalShelves(orders, additionalShelves)
	return nil
}

// groupAdditionalShelves группирует дополнительные стеллажи для каждого продукта.
func groupAdditionalShelves(orders map[string][]model.Order, additionalShelves map[int64][]string) {
	for _, v := range orders {
		for i := range v {
			v[i].ShelveNames = append(v[i].ShelveNames, additionalShelves[v[i].ProductID][1:]...)
		}
	}

	printOrders(orders)
}

// printOrders выводит информацию о заказах.
func printOrders(orders map[string][]model.Order) {
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
