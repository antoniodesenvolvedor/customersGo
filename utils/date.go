package utils

import (
	"fmt"
	"strings"
)

func ConvertDateStringToISODateFormat(date string) (string, error) {
	result := strings.Split(date, "/")
	if len(result) < 3 {
		return "", fmt.Errorf("não foi possível converter a string %s", date)
	}
	day := result[0]
	month := result[1]
	year := result[2]

	return fmt.Sprintf("%s-%s-%s", year, month, day), nil

}
