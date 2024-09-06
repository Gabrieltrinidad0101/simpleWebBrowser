package utils

import (
	"strconv"
	"strings"
)

func percentage(value string, container float32) (float32, error) {
	value = strings.ReplaceAll(value, "%", "")
	number, err := strconv.ParseFloat(value, 32)
	return float32(number) * container / 100, err
}

func NumberDefault(value string, container float32, defaultValue float32) float32 {
	if strings.Contains(value, "px") {
		value = strings.ReplaceAll(value, "px", "")
		number, err := strconv.ParseFloat(value, 32)
		if err != nil {
			return defaultValue
		}
		return float32(number)
	}

	if strings.Contains(value, "%") {
		number, err := percentage(value, container)
		if err != nil {
			return defaultValue
		}
		return float32(number)
	}

	return defaultValue

}
