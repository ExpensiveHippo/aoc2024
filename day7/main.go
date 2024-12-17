package main

import (
	"os"
	"fmt"
	"strings"
	"strconv"
)

func main() {
	f := func(c rune) bool {
		return c == '\n'
	}
	file, _ := os.ReadFile("day7/input.txt")
	res, res2 := 0, 0
	for _, line := range strings.FieldsFunc(string(file), f) {
		data := strings.Split(strings.TrimSpace(line), ":")
		target, _ := strconv.Atoi(data[0])
		numbers := []int{}
		for _, n := range strings.Split(strings.TrimSpace(data[1]), " ") {
			num, _ := strconv.Atoi(n) 
			numbers = append(numbers, num)
		}
		res += run(target, numbers, false)
		res2 += run(target, numbers, true)
	}
	fmt.Println("Part 1: ", res)
	fmt.Println("Part 2: ", res2)
}

func run(target int, values []int, concat bool) int {
	if len(values) == 0 { return 0 }
	if len(values) == 1 {
		if values[0] == target {
			return target
		}
		return 0
	}

	if r := run(target, 
	append([]int{values[0] + values[1]}, values[2:]...),
	concat);
	r != 0 { 
		return target 
	}

	if r := run(target, 
	append([]int{values[0] * values[1]}, values[2:]...),
	concat); 
	r != 0 { 
		return target
	}

	if concat {
		newVal, _ := strconv.Atoi(strconv.Itoa(values[0]) + strconv.Itoa(values[1]))
		if r := run(target, 
		append([]int{ newVal }, values[2:]...),
		concat);
		r != 0 {
			return target
		}
	}
	
	return 0
}

