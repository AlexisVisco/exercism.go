package reverse

import (
	"strings"
)

func String(str string) string {
	var strBuilder strings.Builder
	if len(str) > 0 {
		runes := []rune(str)
		for e := range runes {
			strBuilder.WriteRune(runes[(len(runes) - 1) - e])
		}
	}
	return strBuilder.String()
}
