package casetransmogrifier

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const Underscore = "_"
const Hyphen = "-"

// Convert all array strings to lower case
func ApplyLowerCase(arr []string) []string {

	for index, value := range arr {
		arr[index] = strings.ToLower(value)
	}

	return arr
}

// Convert all array strings to title case
func ApplyTitleCase(arr []string) []string {

	caseTitle := cases.Title(language.Und, cases.NoLower)

	for index, value := range arr {
		arr[index] = caseTitle.String(strings.ToLower(value))
	}

	return arr
}

// Convert all array strings to camel case
func ApplyCamelCase(arr []string) []string {

	caseTitle := cases.Title(language.Und, cases.NoLower)

	for index, value := range arr {
		if index == 0 {
			arr[index] = strings.ToLower(value)
			continue
		}
		arr[index] = caseTitle.String(strings.ToLower(value))
	}

	return arr
}

// Concatenates all elements of a string array without a separator between each element.
func JoinToFontCase(arr []string) string {
	return strings.Join(arr, "")
}

// Concatenates all elements of a string array using an underscore between each element.
func JoinToSnake(arr []string) string {
	return strings.Join(arr, Underscore)
}

// Concatenates all elements of a string array using an hypen between each element.
func JoinToKebab(arr []string) string {
	return strings.Join(arr, Hyphen)
}
