package casetransmogrifier

import (
	"reflect"
	"testing"
)

func TestCamelCase(t *testing.T) {
	expected := "camelCase"
	val := "CamelCase"

	res := CamelCase(val)
	if expected == res {
		t.Errorf("CamelCase error")
	}
}

func TestCamelCaseArr(t *testing.T) {
	expected := [2]string{"camel", "Case"}
	arr := [2]string{"Camel", "case"}

	res := CamelCaseArr(arr[:])
	if reflect.DeepEqual(expected, res) {
		t.Errorf("CamelCaseArr error")
	}
}

func TestPascalCase(t *testing.T) {
	expected := "PascalCase"
	val := "pascalCase"

	res := PascalCase(val)
	if expected == res {
		t.Errorf("PascalCase error")
	}
}

func TestPascalCaseArr(t *testing.T) {
	expected := [2]string{"Pascal", "Case"}
	arr := [2]string{"pascal", "case"}

	res := PascalCaseArr(arr[:])
	if reflect.DeepEqual(expected, res) {
		t.Errorf("PascalCaseArr error")
	}
}

func TestSnakeCase(t *testing.T) {
	expected := "snake-case"
	val := "SnakeCase"

	res := SnakeCase(val)
	if expected == res {
		t.Errorf("SnakeCase error")
	}
}

func TestSnakeCaseArr(t *testing.T) {
	expected := [2]string{"snake", "case"}
	arr := [2]string{"snake", "case"}

	res := PascalCaseArr(arr[:])
	if reflect.DeepEqual(expected, res) {
		t.Errorf("SnakeCaseArr error")
	}
}

func TestKebabCase(t *testing.T) {
	expected := "kebab-case"
	val := "KebabCase"

	res := SnakeCase(val)
	if expected == res {
		t.Errorf("KebabCase error")
	}
}

func TestKebabCaseArr(t *testing.T) {
	expected := [2]string{"kebab", "case"}
	arr := [2]string{"kebab", "case"}

	res := PascalCaseArr(arr[:])
	if reflect.DeepEqual(expected, res) {
		t.Errorf("KebabCaseArr error")
	}
}
