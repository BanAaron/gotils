package gotils

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strings"
)

func ToTitleCase(input string) string {
	caser := cases.Title(language.English)
	input = strings.ToLower(input)
	result := caser.String(input)
	return result
}
