package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func ConvertPrice(priceString string) (float64, error) {
	if priceString == "" {
		return 0, fmt.Errorf("empty price")
	}

	clean := strings.ReplaceAll(priceString, "R$", "")
	clean = strings.ReplaceAll(clean, " ", "")
	clean = strings.ReplaceAll(clean, ".", "")

	clean = strings.ReplaceAll(clean, ",", ".")

	price, err := strconv.ParseFloat(clean, 64)

	if err != nil {
		return 0, fmt.Errorf("error trying to convert price %s: %v", priceString, err)
	}

	return price, nil
}
