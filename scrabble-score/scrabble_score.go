package scrabble

import (
	"strings"
	"fmt"
)

const (
	PAL_1 string = "AEIOULNRST"
	PAL_2 string = "DG"
	PAL_3 string = "BCMP"
	PAL_4 string = "FHVWY"
	PAL_5 string = "K"
	PAL_6 string = "JX"
	PAL_7 string = "QZ"
)

var points = CreateMap()

func CreateMap() map[byte]int {
	var a = make(map[byte]int)
	SetFor(a, PAL_1, 1)
	SetFor(a, PAL_2, 2)
	SetFor(a, PAL_3, 3)
	SetFor(a, PAL_4, 4)
	SetFor(a, PAL_5, 5)
	SetFor(a, PAL_6, 8)
	SetFor(a, PAL_7, 10)
	return a
}

func SetFor(points map[byte]int, str string, point int) {
	for e := range str {
		points[str[e]] = point
	}
}

func Score(word string) int {
	fmt.Print(points)
	score := 0
	word = strings.ToUpper(word)
	for e := range word {
		score += points[word[e]]
	}
	return score
}