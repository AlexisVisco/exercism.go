package twofer

import "fmt"

func ShareWith(str string) string {
	if str == "" {
		return "One for you, one for me."
	}
	return fmt.Sprintf("One for %s, one for me.", str)
}
