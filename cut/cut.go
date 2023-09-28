package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// определяем флаги
	fields := flag.String("f", "", "выбрать поля")
	delimiter := flag.String("d", "\t", "использовать разделитель")
	separated := flag.Bool("s", false, "только строки с разделителем")
	flag.Parse()

	// обрабатываем входные данные
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if *separated && !strings.Contains(line, *delimiter) {
			continue
		}
		parts := strings.Split(line, *delimiter)
		if *fields == "" {
			fmt.Println(line)
		} else {
			indices := strings.Split(*fields, ",")
			for _, indexStr := range indices {
				index := strings.TrimSpace(indexStr)
				if index == "" {
					continue
				}
				i, err := strconv.Atoi(index)
				if err != nil || i < 1 || i > len(parts) {
					continue
				}
				fmt.Print(parts[i-1])
				if i < len(parts) {
					fmt.Print(*delimiter)
				}
			}
			fmt.Println()
		}
	}
}
