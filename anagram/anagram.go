package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	m := []string{"Пятак", "Пятка", "Тяпка", "Листок", "Слиток", "Листок", "Столик"}
	s := searchAnagram(m)
	fmt.Printf("%T, %v", s, s)
}

func searchAnagram(s []string) map[string][]string {
	anagramSet := make(map[string][]string)
	var lenSlice int
	for i := 0; i < len(s); i += lenSlice {
		sliceSet := make([]string, 0)
		lowStr := strings.ToLower(s[i])
		sliceSet = append(sliceSet, lowStr)
		splitedString := strings.Split(lowStr, "")
		sort.Strings(splitedString)

		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				continue
			} else {
				var splittedString2 []string
				lowStrA := strings.ToLower(s[j])
				splittedString2 = strings.Split(lowStrA, "")
				sort.Strings(splittedString2)

				if CompareSlice(splitedString, splittedString2) == true {
					sliceSet = append(sliceSet, lowStrA)
				}
			}
		}
		sort.Strings(sliceSet)
		lenSlice = len(sliceSet)
		if lenSlice > 1 {
			key := sliceSet[0]
			anagramSet[key] = sliceSet
		}
	}
	if len(anagramSet) == 0 {
		return nil
	} else {
		return anagramSet
	}

}

func CompareSlice(slice1, slice2 []string) bool {
	for letter := 0; letter < len(slice1); letter++ {
		if slice1[letter] != slice2[letter] {
			return false
		}
	}
	return true
}
