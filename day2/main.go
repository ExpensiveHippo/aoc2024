package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	safeCount := countSafeReports()
	safeCount2 := countSafeReports2()
	fmt.Println("Safe:", safeCount)
	fmt.Println("Safe 2:", safeCount2)
}

/*
Return number of successful reports. Successful if:
- Always ascending/descending 
- Difference between 2 consecutive level is less than or equal to 3
*/
func countSafeReports() (count int) {
	file, err := os.Open("./day2/input.txt")

	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var levelArr []int
		report := scanner.Text()
		data := strings.Split(report, " ")

		for _, i := range data {
			level, _ := strconv.Atoi(i)
			levelArr = append(levelArr, level)
		}

		if isValid(levelArr) {
			count++
		}
	}
	return
}

/*
Add a dampener which allows a level to be removed from the test case to make it successful
Follow same criteria as previous stage
*/
func countSafeReports2() (count int) {
	file, err := os.Open("./day2/input.txt")

	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var levelArr []int
		report := scanner.Text()
		data := strings.Split(report, " ")

		for _, i := range data {
			level, _ := strconv.Atoi(i)
			levelArr = append(levelArr, level)
		}

		if isValid(levelArr) {
			count++
		} else {
			safe := checkDampen(levelArr)
			if safe {
				count ++
			}
		}
	}
	return
}

// consider successful if removing one level makes the test successful
func checkDampen(levelArr []int) bool {
	for i := range levelArr {
		temp := make([]int, len(levelArr)-1)
		k := 0
		for j, jVal := range levelArr {
			if i != j {
				temp[k] = jVal
				k++
			}
		}

		if isValid(temp) {
			return true
		}
	}
	return false
}

func isValid(levelArr []int) bool {
	return (isAsc(levelArr) || isDesc(levelArr)) && !hasDuplicates(levelArr) && lowDifference(levelArr)
}

func isAsc(levelArr []int) bool {
	asc := append([]int{}, levelArr...)
	slices.Sort(asc)
	return slices.Equal(levelArr, asc)
}

func isDesc(levelArr []int) bool {
	desc := append([]int{}, levelArr...)
	slices.SortStableFunc(desc, func(a, b int) int {
		return b - a
	})
	return slices.Equal(levelArr, desc)
}

func hasDuplicates(levelArr []int) bool {
	freq := make(map[int]int)
	for _, i := range levelArr {
		if freq[i] != 0 {
			return true
		} else {
			freq[i]++
		}
	}
	return false
}

func lowDifference(levelArr []int) bool {
	safe := true
	for i := 0; i < len(levelArr)-1; i++ {
		diff := levelArr[i] - levelArr[i+1]
		if diff > 3 || diff < -3 {
			safe = false
			break
		}
	}
	return safe
}
