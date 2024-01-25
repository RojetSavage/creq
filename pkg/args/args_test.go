package args

import (
	"testing"
)

func TestIsFlag(t *testing.T) {
	var tests = []struct {
		input    string
		expected bool
	}{
		{"--json", true},
		{"--connect-timeout", true},
		{"--c*onnect-timeout", false},
		{"--743298", false},
		{"--", false},
		{"", false},
		{"-- ", false},
		{"----", false},
		{"--json!", false},
		{"--j", false},
	}
	for _, test := range tests {
		if b, _ := isFlag(test.input); b != test.expected {
			t.Error("Failed")
		}
	}
}
func TestIsShort(t *testing.T) {
	var tests = []struct {
		input    string
		expected bool
	}{
		{"-p", true},
		{"-I", true},
		{"- ", false},
		{"--", false},
		{"-&", false},
		{"  ", false},
		{"-4", false},
		{"-+", false},
	}

	for _, test := range tests {
		if b, _ := isShort(test.input); b != test.expected {
			t.Error("Failed")
		}
	}
}
