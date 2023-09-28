package main

import (
	"reflect"
	"testing"
)

func TestAnagrams(t *testing.T) {
	testCases := []struct {
		input    []string
		expected map[string][]string
	}{
		{
			input: []string{"Пятак", "Пятка", "Тяпка", "Листок", "Слиток", "Листок", "Столик"},
			expected: map[string][]string{
				"пятак":  {"пятак", "пятка", "тяпка"},
				"листок": {"листок", "слиток", "столик"},
			},
		},
		{
			input:    []string{"яблоко", "колба", "машина"},
			expected: nil,
		},
	}
	for _, test := range testCases {
		t.Run("", func(t *testing.T) {
			result := searchAnagram(test.input)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("Ожидается %v, получили %v", test.expected, result)
			}
		})
	}

}
