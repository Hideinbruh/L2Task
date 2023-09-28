package main

import "testing"

func TestUnpack(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"a4bc2d5e", "aaaabccddddde"},
		{"abcd", "abcd"},
		{"45", ""},
		{"", ""},
		{"qwe\\4\\5", "qwe45"},
		{"qwe\\45", "qwe44444"},
		{"qwe\\\\5", "qwe\\\\\\\\\\"},
	}

	for _, testCase := range testCases {
		t.Run(testCase.input, func(t *testing.T) {
			result := unpackString(testCase.input)
			if result != testCase.expected {
				t.Errorf("For data %s got result %s\n expected result: %s", testCase.input, result, testCase.expected)
			}
		})
	}
}
