package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	sum := sumMull()
	fmt.Println("Sum: ", sum)
}

func sumMull() (sum int) {
	file, err := os.Open("./day3/input.txt")

	if err != nil {
		panic(err)
	}
	defer file.Close()

	// check for mul(int1, int2) or don't() or do()
	regex := regexp.MustCompile(`mul\(\d+,\d+\)|don't\(\)|do\(\)`)

	scanner := bufio.NewScanner(file)

	doCheck := true

	for scanner.Scan() {
		input := scanner.Text()
		commands := regex.FindAllString(input, -1)
		for _, c := range commands {
			if c == "don't()" {
				doCheck = false
			} else if c == "do()" {
				doCheck = true
				continue
			}

			if !doCheck {
				continue
			}

			args := strings.Split(c, ",")
			x, err := strconv.Atoi(strings.Split(args[0], "(")[1])

			if err != nil {
				panic(err)
			}

			y, err := strconv.Atoi(strings.Split(args[1], ")")[0])

			if err != nil {
				panic(err)
			}
			sum += x * y
		}
	}
	return
}
