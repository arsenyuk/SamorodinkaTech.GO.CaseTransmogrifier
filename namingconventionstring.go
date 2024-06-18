package casetransmogrifier

// Camel Case (camelCase)
func CamelCase(val string) string {
	return CamelCaseArr(ParseCase(val))
}

// Camel Case (camelCase)
func CamelCaseArr(val []string) string {
	return JoinToFontCase(ApplyCamelCase(val))
}

// Pascal Case (PascalCase)
func PascalCase(val string) string {
	return PascalCaseArr(ParseCase(val))
}

// Pascal Case (PascalCase)
func PascalCaseArr(val []string) string {
	return JoinToFontCase(ApplyTitleCase(val))
}

// Snake Case (snake_case)
func SnakeCase(val string) string {
	return SnakeCaseArr(ParseCase(val))
}

// Snake Case (snake_case)
func SnakeCaseArr(val []string) string {
	return JoinToSnake(ApplyLowerCase(val))
}

// Kebab Case (kebab-case)
func KebabCase(val string) string {
	return KebabCaseArr(ParseCase(val))
}

// Kebab Case (kebab-case)
func KebabCaseArr(val []string) string {
	return JoinToKebab(ApplyLowerCase(val))
}
