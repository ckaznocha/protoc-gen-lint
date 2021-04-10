package linter

import (
	"testing"
)

func Test_isCamelCase(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{"snake case", args{"hello_world"}, false},
		{"screaming snake case", args{"HELLO_WORLD"}, false},
		{"dromedary case", args{"helloWorld"}, false},
		{"lower case", args{"helloworld"}, false},
		{"screeming", args{"HELLOWORLD"}, false},
		{"Pascal case", args{"HelloWorld"}, true},
		{"dromedary case with first a signel leading alpha", args{"h264"}, false},
		{"Pascal case with a single leading alpha", args{"H264"}, true},
		{"Pascal case with some numbers after a leading alpha", args{"H264Encoded"}, true},
		{"screeming with some numbers after a leading alpha", args{"H264ENCODED"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isCamelCase(tt.args.s); got != tt.want {
				t.Errorf("isCamelCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isLowerUnderscore(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{"snake case", args{"hello_world"}, true},
		{"screaming snake case", args{"HELLO_WORLD"}, false},
		{"dromedary case", args{"helloWorld"}, false},
		{"lower case", args{"helloworld"}, true},
		{"screaming", args{"HELLOWORLD"}, false},
		{"Pascal case ", args{"HelloWorld"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isLowerUnderscore(tt.args.s); got != tt.want {
				t.Errorf("isLowerUnderscore(() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isUpperUnderscore(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{"snake case", args{"hello_world"}, false},
		{"screaming snake case", args{"HELLO_WORLD"}, true},
		{"dromedary case", args{"helloWorld"}, false},
		{"lower case", args{"helloworld"}, false},
		{"screaming", args{"HELLOWORLD"}, true},
		{"Pascal case ", args{"HelloWorld"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isUpperUnderscore(tt.args.s); got != tt.want {
				t.Errorf("sUpperUnderscore(() = %v, want %v", got, tt.want)
			}
		})
	}
}
