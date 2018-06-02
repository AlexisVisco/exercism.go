package isogram

import "strings"

func IsIsogram(input string) bool {
	letters := make(map[byte]int)
	input = strings.ToLower(input)
	for e := range input {
		if input[e] == ' ' || input[e] == '-' {
			continue
		}
		if _, ok := letters[input[e]]; ok {
			return false
		}
		letters[input[e]] = 1
	}
	return true
}