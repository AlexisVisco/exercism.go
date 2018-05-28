package acronym

import (
	"strings"
	"bytes"
)

func Abbreviate(s string) string {
	var arr = strings.Split(s, " ")
	var buffer bytes.Buffer
	for e := range arr {
		var cutted = strings.Split(arr[e], "-");
		for e := range cutted {
			buffer.WriteByte(cutted[e][0])
		}
	}
	return strings.ToUpper(buffer.String())
}
