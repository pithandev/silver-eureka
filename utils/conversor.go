package utils

import (
	"strconv"
	"strings"
	"unicode"
)

func ConvertPrice(priceString string) (float64, error) {
	if priceString == "" {
		return 0, nil
	}

	var builder strings.Builder
	for _, char := range priceString {
		if unicode.IsDigit(char) || char == '.' || char == ',' {
			builder.WriteRune(char)
		}
	}

	clean := builder.String()

	if clean == "" {
		return 0, nil
	}

	if strings.Contains(clean, ",") {
		if strings.Contains(clean, ".") {
			clean = strings.ReplaceAll(clean, ".", "")
		}
		clean = strings.ReplaceAll(clean, ",", ".")
	}

	return strconv.ParseFloat(clean, 64)
}
