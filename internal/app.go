package internal

import (
	"context"
	"fmt"
	db "github.com/ent1k1377/testovoe_20_03_24/db/sqlc"
	"sort"
	"strings"
)

type Orders struct {
	productId   int64
	productName string
	orderId     int64
	quantity    int32
	shelveNames []string
}

func Start(sqlStore *db.SQLStore, orderNumbers []int64) error {
	ctx := context.Background()
	orderInfo, err := sqlStore.Queries.GetOrderInfo(ctx, orderNumbers)
	if err != nil {
		return err
	}

	q := make(map[int64][]string)
	m := make(map[string][]Orders)
	keys := make([]string, 0)
	for _, e := range orderInfo {
		if !e.ShelveIsPrimary {
			q[e.ProductID] = append(q[e.ProductID], e.ShelveName)
			continue
		}

		q[e.ProductID] = []string{e.ShelveName}

		m[e.ShelveName] = append(m[e.ShelveName], Orders{
			productId:   e.ProductID,
			productName: e.ProductName,
			orderId:     e.OrderID,
			quantity:    e.Quantity,
			shelveNames: make([]string, 0),
		})
	}

	for _, v := range m {
		for i := range v {
			v[i].shelveNames = append(v[i].shelveNames, q[v[i].productId][1:]...)
		}
	}

	for k := range m {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	writeOrders(m, keys)
	return nil
}

func writeOrders(m map[string][]Orders, sortedKeys []string) {
	for _, e := range sortedKeys {
		fmt.Printf("===Стеллаж %s\n", e)
		v := m[e]
		for _, o := range v {
			fmt.Printf("%s (id=%d)\nзаказ %d, %d шт\n",
				o.productName,
				o.productId,
				o.orderId,
				o.quantity)
			if len(o.shelveNames) != 0 {
				fmt.Printf("доп стеллаж: %s\n", strings.Join(o.shelveNames, ","))
			}
			fmt.Println()
		}
	}
}
