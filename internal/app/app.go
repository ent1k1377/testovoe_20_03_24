package app

import (
	"fmt"
	"github.com/ent1k1377/testovoe_20_03_24/internal/repository"
	"github.com/ent1k1377/testovoe_20_03_24/internal/repository/postgres"
	"github.com/ent1k1377/testovoe_20_03_24/internal/util"
	"log"
	"os"
	"sort"
	"strings"
)

// order представляет информацию о заказе.
type order struct {
	OrderID     int64
	ProductID   int64
	ProductName string
	Quantity    int32
	ShelveNames []string
}

type Product struct {
	name     string
	quantity int32
}

type GetOrderInfoRow struct {
	ProductName     string `json:"product_name"`
	ProductID       int64  `json:"product_id"`
	OrderID         int64  `json:"order_id"`
	Quantity        int32  `json:"quantity"`
	ShelveName      string `json:"shelve_name"`
	ShelveIsPrimary bool   `json:"shelve_is_primary"`
}

// StartProcessingOrders начинает обработку заказов.
func StartProcessingOrders(storage repository.Storage, orderNumbers []string) error {
	// Преобразование аргументов командной строки в числа
	ordersId, err := util.ConvertStringsToIntegers(os.Args[1:])
	if err != nil {
		log.Fatal("Incorrect transmitted data", err)
	}

	var ordersInfo []GetOrderInfoRow

	q := make(map[int64]Product)

	w := make(map[int64]int32)
	for _, id := range ordersId {

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
			w[v.ProductID] = v.Quantity
		}
		for _, v := range products {
			q[v.ID] = Product{
				name:     v.Name,
				quantity: w[v.ID],
			}
		}
		shelves := map[int64]string{
			1: "А",
			2: "Б",
			3: "В",
			4: "З",
			5: "Ж",
		}

		for _, v := range productsShelves {
			ordersInfo = append(ordersInfo, GetOrderInfoRow{
				ProductName:     q[v.ProductID].name,
				ProductID:       v.ProductID,
				OrderID:         id,
				Quantity:        q[v.ProductID].quantity,
				ShelveName:      shelves[v.ShelvesID],
				ShelveIsPrimary: v.IsPrimary,
			})
		}
	}

	fmt.Printf("=+=+=+=\nСтраница сборки заказов %s\n\n", strings.Join(orderNumbers, ","))
	return processOrders(ordersInfo)
}

// processOrders обрабатывает информацию о заказах.
func processOrders(ordersInfo []GetOrderInfoRow) error {
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
			OrderID:     o.OrderID,
			ProductID:   o.ProductID,
			ProductName: o.ProductName,
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
