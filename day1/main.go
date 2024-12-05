package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"slices"
)

func main() {
	i := findDistance1()
	j := findDistance2()
	fmt.Println("Stage 1: ", i);
	fmt.Println("Stage 2: ", j)
}

// Stage 1: Sum difference between each pair from first and second list (sorted)
func findDistance1() (totalDistance int) {
	var list1 []int
	var list2 []int

	totalDistance = 0
	file, err := os.Open("./day1/input.txt")

	if err != nil {
		fmt.Println(err.Error())
		return -1
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		pair := strings.Split(scanner.Text(), "   ")
		x, _ := strconv.Atoi(pair[0])
		y, _ := strconv.Atoi(pair[1])
		list1 = append(list1, x)
		list2 = append(list2, y)
	}

	slices.SortStableFunc(list1, func(a, b int) int {
		return a - b
	})

	slices.SortStableFunc(list2, func(a, b int) int {
		return a - b
	})

	for i := 0; i < len(list1); i++ {
		distance := list1[i] - list2[i]
		if (distance < 0) {
			distance *= -1
		}
		totalDistance += distance
	}
	
	return
}

// Stage 2: Sum the product of every digit in the first list with its frequency in the second list
func findDistance2() (totalDistance int) {
	var list1 []int
	freq2 := make(map[int]int)

	totalDistance = 0
	file, err := os.Open("./day1/input.txt")

	if err != nil {
		fmt.Println(err.Error())
		return -1
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		pair := strings.Split(scanner.Text(), "   ")
		x, _ := strconv.Atoi(pair[0])
		y, _ := strconv.Atoi(pair[1])
		list1 = append(list1, x)
		freq2[y] += 1
	}

	for i := 0; i < len(list1); i++ {
		distance := list1[i] * freq2[list1[i]]
		if (distance < 0) {
			distance *= -1
		}
		totalDistance += distance
	}
	
	return
}

