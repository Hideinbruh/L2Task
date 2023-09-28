package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	afterLines := flag.Int("A", 0, "Печатать +N строк после совпадения")
	beforeLines := flag.Int("B", 0, "Печатать +N строк до совпадения")
	contextLines := flag.Int("C", 0, "Печатать ±N строк вокруг совпадения")
	countOnly := flag.Bool("c", false, "Печатать только количество совпадений")
	ignoreCase := flag.Bool("i", false, "Игнорировать регистр при поиске")
	invertMatch := flag.Bool("v", false, "Исключить строки с совпадениями")
	fixedString := flag.Bool("F", false, "Искать точное совпадение со строкой, не как паттерн")
	lineNumbers := flag.Bool("n", false, "Печатать номера строк")

	flag.Parse()

	args := flag.Args()
	pattern := args[0]
	fileName := args[1]

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	wordRegex := regexp.MustCompile(pattern)
	scanner := bufio.NewScanner(file)
	data := make(map[int]string)
	var found bool
	var lines string
	var linesCount int
	for i := 1; scanner.Scan(); i++ {
		lines = scanner.Text()
		linesCount++
		if *afterLines > 0 {

			if found {
				linesCount = 0
				fmt.Println(lines)
				linesCount++

				if linesCount >= *afterLines {
					linesCount = 0
					found = false
					break

				}
			} else if wordRegex.MatchString(lines) {
				found = true
				fmt.Println(pattern)

			}
		}

		if *beforeLines > 0 {

			data[i] = lines

			if found {
				for j := i - 1; j >= i-*beforeLines && j > 0; j-- {
					fmt.Println(data[j-1])

				}
				found = false
				break
			} else if wordRegex.MatchString(lines) {
				found = true
				fmt.Println(pattern)
			}
		}
		if *contextLines > 0 {
			data[i] = lines

			if found {
				linesCount = 0
				fmt.Println(lines)
				linesCount++

				if linesCount >= *contextLines {
					for j := i - linesCount; j > i-linesCount-*contextLines && j > 0; j-- {
						fmt.Println(data[j-1])
					}
					found = false
					linesCount = 0
					break
				}
			} else if wordRegex.MatchString(lines) {
				found = true
				fmt.Println(pattern)
			}
		}
		if *ignoreCase {
			lowerStr := strings.ToLower(lines)
			lowerPattern := strings.ToLower(pattern)
			wordPatternForCaseI := regexp.MustCompile(lowerPattern)

			if found {
				fmt.Printf("Игнорирование регистра при поиске. Значение %s найдено на строке: %d", pattern, linesCount-1)
				found = false
				break

			} else if wordPatternForCaseI.MatchString(lowerStr) {
				found = true
			}
		}
		if *invertMatch {
			if found {
				fmt.Println(lines)
			} else if wordRegex.MatchString(lines) {
				found = true
			} else if wordRegex.MatchString(lines) != true {
				fmt.Println(lines)
			}
		}
		if *fixedString {
			if found {
				fmt.Printf("Точное совпадение со строкой найдено на строке %d, значение: %s", linesCount, pattern)
				found = false
				break

			} else if lines == pattern {
				found = true
				fmt.Println(pattern)
			}

		}
		if *lineNumbers {
			fmt.Println(linesCount, lines)
		}
	}
	if *countOnly {
		fmt.Println("Количество строк:", linesCount)
	}
}