package linter

import "testing"

func TestIsCamelCase(t *testing.T) {
	var stringsToTest = []struct {
		test string
		want bool
	}{
		{"hello_world", false},
		{"HELLO_WORLD", false},
		{"helloWorld", false},
		{"helloworld", false},
		{"HELLOWORLD", false},
		{"HelloWorld", true},
		{"h264", false},
		{"H264", true},
	}

	for _, v := range stringsToTest {
		if got := isCamelCase(v.test); got != v.want {
			t.Errorf("Expected %t, Received %t", v.want, got)
		}
	}
}

func TestIsLowerUnderscore(t *testing.T) {
	var stringsToTest = []struct {
		test string
		want bool
	}{
		{"hello_world", true},
		{"HELLO_WORLD", false},
		{"helloWorld", false},
		{"helloworld", true},
		{"HELLOWORLD", false},
		{"HelloWorld", false},
	}

	for _, v := range stringsToTest {
		if got := isLowerUnderscore(v.test); got != v.want {
			t.Errorf("Expected %t, Received %t", v.want, got)
		}
	}
}

func TestIsUpperUnderscore(t *testing.T) {
	var stringsToTest = []struct {
		test string
		want bool
	}{
		{"hello_world", false},
		{"HELLO_WORLD", true},
		{"helloWorld", false},
		{"helloworld", false},
		{"HELLOWORLD", true},
		{"HelloWorld", false},
	}

	for _, v := range stringsToTest {
		if got := isUpperUnderscore(v.test); got != v.want {
			t.Errorf("Expected %t, Received %t", v.want, got)
		}
	}
}
