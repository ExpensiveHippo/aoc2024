package main 

import (
	"os"
	"fmt"
	"strings"
	"strconv"
)

func main() {
	file, _ := os.ReadFile("day11/input.txt")
	stones := map[int]int{}
	
	for _, i := range strings.Split(strings.TrimSpace(string(file)), " ") {
		stone, _ := strconv.Atoi(i)
		stones[stone]++
	}

	results := map[int][]int{}
	

	for i := 0; i < 75; i++ {
		newStones := map[int]int{}

		for s, count := range stones {
			var a []int
			if _, ok := results[s]; !ok {
				a = transform(s)
				results[s] = a
			} else {
				a = results[s]
			}
			
			for _, b := range a {
				newStones[b] += count
			}
		}

		stones = newStones
	}

	ans := 0

	for _, i := range stones {
		ans += i
	}

	fmt.Println(ans)
}

func transform(s int) []int {
	if s == 0 {
		return []int{ 1 }
	} 

	if a := strconv.Itoa(s); len(a) % 2 == 0 {
		stone1, _ := strconv.Atoi(a[:len(a)/2])
		stone2, _ := strconv.Atoi(a[len(a)/2:])
		return []int{ stone1, stone2 }
	} 

	return []int{ s * 2024 }
}
