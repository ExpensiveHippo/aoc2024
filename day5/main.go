package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
	"strconv"
)

func main() {
	file, _ := os.ReadFile("./day5/input.txt")
	data := strings.Split(strings.ReplaceAll(string(file), "\r\n", "\n"), "\n\n")
	rules := make(map[string][]string)

	for _, rule := range strings.Fields(data[0]) {
		r := strings.Split(rule, "|")
		rules[r[1]] = append(rules[r[1]], r[0])
	}

	compare := func (a, b string) int {
		if slices.Contains(rules[b], a) {
			return -1
		} 

		return 0
	}

	solve := func(checkSorted bool) (sum int) {
		for _, input := range strings.Fields(data[1]) {
			pages := strings.Split(input, ",")

			if slices.IsSortedFunc(pages, compare) == checkSorted {
				slices.SortStableFunc(pages, compare)
				mid, _ := strconv.Atoi(pages[len(pages)/2])
				sum += mid
			}
		}
		return
	}

	fmt.Println("Sum (sorted): ", solve(true))
	fmt.Println("Sum (unsorted): ", solve(false))
}
