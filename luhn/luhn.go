package luhn

import (
	"strings"
	"strconv"
)

func Valid(digits string) bool {
	digitsList := strings.Split(strings.Replace(digits, " ", "", -1), "")
	sum, err := strconv.Atoi(digitsList[len(digitsList)-1])
	parity := 1 - ((len(digitsList) - 1) % 2)
	if err != nil || len(digitsList) == 1 && digitsList[0] == "0"{
		return false
	}
	for i := len(digitsList) - 2; i >= 0; i -- {
		digit, err := strconv.Atoi(digitsList[i])
		if err != nil {
			return false
		}
		if i%2 == parity {
			digit = digit * 2
			if digit > 9 {
				digit = digit - 9
			}
		}
		sum += digit
	}
	return (sum % 10) == 0
}