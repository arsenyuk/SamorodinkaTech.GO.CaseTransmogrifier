package casetransmogrifier

import (
	"reflect"
	"testing"
)

func TestApplyLowerCase(t *testing.T) {
	expected := [2]string{"a", "n"}
	arr := [2]string{"a", "N"}

	res := ApplyLowerCase(arr[:])
	if reflect.DeepEqual(expected, res) {
		t.Errorf("ApplyLowerCase error")
	}
}

func TestApplyTitleCase(t *testing.T) {
	expected := [2]string{"A", "N"}
	arr := [2]string{"a", "N"}

	res := ApplyTitleCase(arr[:])
	if reflect.DeepEqual(expected, res) {
		t.Errorf("ApplyTitleCase error")
	}
}

func TestApplyCamelCase(t *testing.T) {
	expected := [2]string{"a", "n"}
	arr := [2]string{"a", "N"}

	res := ApplyCamelCase(arr[:])
	if reflect.DeepEqual(expected, res) {
		t.Errorf("ApplyCamelCase error")
	}
}

func TestJoinToFontCase(t *testing.T) {
	expected := "FontCase"
	arr := [2]string{"Font", "Case"}

	res := JoinToFontCase(arr[:])
	if res != expected {
		t.Errorf("JoinToFontCase error")
	}
}

func TestJoinToSnake(t *testing.T) {
	expected := "snake_case"
	arr := [2]string{"snake", "case"}

	res := JoinToSnake(arr[:])
	if res != expected {
		t.Errorf("JoinToSnake error")
	}
}

func TestJoinToKebab(t *testing.T) {
	expected := "kebab-case"
	arr := [2]string{"kebab", "case"}

	res := JoinToKebab(arr[:])
	if res != expected {
		t.Errorf("JoinToKebab error")
	}
}
