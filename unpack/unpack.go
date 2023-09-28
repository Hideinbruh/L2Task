package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func unpackString(data string) string {
	var result strings.Builder
	var escape bool
	for i := 0; i < len(data); i++ {
		if escape == true {
			result.WriteString(string(data[i]))
			escape = false
		} else if data[i] == '\\' {
			escape = true
		} else if unicode.IsDigit(rune(data[i])) {
			count := ""
			for j := i; j < len(data); j++ {
				if unicode.IsDigit(rune(data[j])) {
					count += string(data[j])
				} else {
					break
				}
			}
			num, err := strconv.Atoi(count)
			if num > 0 {
				if i > 0 {
					if err != nil {
						fmt.Printf("Error converting %v", err)
					}
					result.WriteString(strings.Repeat(string(data[i-1]), num-1))
				}
			}
			i += len(count) - 1
		} else if unicode.IsLetter(rune(data[i])) {
			result.WriteString(string(data[i]))
		}
	}
	return result.String()
}

func main() {
	testCases := []string{"a4bc2d5e", "abcd", "45", "", "qwe\\4\\5", "qwe\\45", "qwe\\\\5"}
	for _, testCase := range testCases {
		result := unpackString(testCase)
		fmt.Printf("%s => %s\n", testCase, result)
	}
}
