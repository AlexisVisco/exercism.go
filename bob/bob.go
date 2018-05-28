package bob

import (
	"strings"
	"regexp"
)


func Hey(remark string) string {

	var upperValidation, _ = regexp.Compile("[A-Z]+")
	var blank, _ = regexp.Compile("^[ \r\t\n]*$")
	var question, _ = regexp.Compile("\\?([ \r\t\n]*)$")

	if question.MatchString(remark) {
		if strings.ToUpper(remark) == remark && upperValidation.MatchString(remark) {
			return "Calm down, I know what I'm doing!"
		}
		return "Sure."
	}
	if strings.ToUpper(remark) == remark && upperValidation.MatchString(remark) {
		return "Whoa, chill out!"
	}
	if blank.MatchString(remark) {
		return "Fine. Be that way!"
	}
	return "Whatever."
}
