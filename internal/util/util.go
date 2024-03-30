package util

import (
	"errors"
	"strconv"
)

func ConvertStringsToIntegers(data []string) ([]int64, error) {
	converted := make([]int64, len(data))
	for i, _ := range data {
		elem, err := strconv.Atoi(data[i])
		if err != nil {
			return nil, err
		}
		converted[i] = int64(elem)
	}

	if cap(converted) == 0 {
		return nil, errors.New("empty slice")
	}

	return converted, nil
}
