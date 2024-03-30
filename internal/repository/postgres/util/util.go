package util

import (
	"fmt"
	"strings"
)

func GetQuery(query string, length int) string {
	placeholders := make([]string, length)
	for i := range placeholders {
		placeholders[i] = fmt.Sprintf("$%d", i+1)
	}
	placeholder := strings.Join(placeholders, ", ")

	return fmt.Sprintf(query, placeholder)
}

func ConvertInt64SliceToInterfaceSlice(slice []int64) []interface{} {
	interfaceSlice := make([]interface{}, len(slice))
	for i, v := range slice {
		interfaceSlice[i] = v
	}

	return interfaceSlice
}
