package utils

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func AngularString(str string) string {
	//Remove - seperation
	words := strings.Split(str, "-")
	titleCaser := cases.Title(language.Make("en"))
	var result string

	for i, word := range words {
		words[i] = titleCaser.String(word)
	}

	for _, k := range words {
		result += k
	}

	return result
}
