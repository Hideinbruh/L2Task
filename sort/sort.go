package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode"
)

var (
	filePath  string
	column    int
	isNumeric bool
	isReverse bool
	isUnique  bool
	isMonth   bool
	ignoreBlk bool
	checkSort bool
	isSuffix  bool
)

func init() {
	flag.StringVar(&filePath, "f", "", "File path")
	flag.IntVar(&column, "k", 0, "Column to sort on")
	flag.BoolVar(&isNumeric, "n", false, "Sort by numeric value")
	flag.BoolVar(&isReverse, "r", false, "Sort in reverse order")
	flag.BoolVar(&isUnique, "u", false, "Remove duplicate lines")
	flag.BoolVar(&isMonth, "M", false, "Sort by month name")
	flag.BoolVar(&ignoreBlk, "b", false, "Ignore trailing blanks")
	flag.BoolVar(&checkSort, "c", false, "Check if data is sorted")
	flag.BoolVar(&isSuffix, "h", false, "Sort by numeric value with suffixes")
}

func main() {
	flag.Parse()

	if filePath == "" {
		fmt.Println("File path not provided")
		os.Exit(1)
	}

	lines, err := readLines(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	if checkSort {
		if isSorted(lines) {
			fmt.Println("Data is already sorted")
			os.Exit(0)
		} else {
			fmt.Println("Data is not sorted")
			os.Exit(1)
		}
	}

	if isMonth {
		sort.Slice(lines, func(i, j int) bool {
			return monthToNumber(lines[i]) < monthToNumber(lines[j])
		})
	} else {
		sort.Slice(lines, func(i, j int) bool {
			return less(lines[i], lines[j])
		})
	}

	if isUnique {
		lines = unique(lines)
	}

	if isReverse {
		reverse(lines)
	}

	writeLines(filePath, lines)
}

func readLines(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func writeLines(filePath string, lines []string) error {
	data := []byte(strings.Join(lines, "\n"))
	return ioutil.WriteFile(filePath, data, 0644)
}

func less(a, b string) bool {
	if column > 0 {
		fieldsA := strings.Fields(a)
		fieldsB := strings.Fields(b)

		if len(fieldsA) < column || len(fieldsB) < column {
			return a < b
		}

		a = fieldsA[column-1]
		b = fieldsB[column-1]
	}

	if ignoreBlk {
		a = strings.TrimSpace(a)
		b = strings.TrimSpace(b)
	}

	if isNumeric {
		numA, errA := strconv.Atoi(a)
		numB, errB := strconv.Atoi(b)

		if errA == nil && errB == nil {
			if isSuffix {
				numA = removeSuffix(numA, a)
				numB = removeSuffix(numB, b)
			}
			return numA < numB
		}
	}

	if isMonth {
		return monthToNumber(a) < monthToNumber(b)
	}

	return a < b
}

func reverse(lines []string) {
	for i, j := 0, len(lines)-1; i < j; i, j = i+1, j-1 {
		lines[i], lines[j] = lines[j], lines[i]
	}
}

func unique(lines []string) []string {
	set := make(map[string]bool)
	var uniqueLines []string

	for _, line := range lines {
		if _, ok := set[line]; !ok {
			set[line] = true
			uniqueLines = append(uniqueLines, line)
		}
	}

	return uniqueLines
}

func isSorted(lines []string) bool {
	for i := 1; i < len(lines); i++ {
		if less(lines[i], lines[i-1]) {
			return false
		}
	}
	return true
}

func monthToNumber(month string) int {
	t, _ := time.Parse("Jan", month)
	return int(t.Month())
}

func removeSuffix(num int, str string) int {
	suffix := str[strings.IndexFunc(str, func(r rune) bool { return !unicode.IsDigit(r) })]
	if suffix == 'k' {
		num *= 1000
	} else if suffix == 'M' {
		num *= 1000000
	} else if suffix == 'G' {
		num *= 1000000000
	}
	return num
}
