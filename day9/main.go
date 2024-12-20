package main

import (
	"os"
	"fmt"
	"strings"
	"strconv"
)

type File struct {
	id, index, size int
}

func main() {
	fmt.Println("Part 1: ", run(false))
	fmt.Println("Part 2: ", run(true))
}

func run(part2 bool)(sum int) {
	var files, spaces []File
	var final []int
	a, b, index := -1, 0, 0

	input, _ := os.ReadFile("day9/input.txt")
	for i, n := range strings.Split(strings.TrimSpace(string(input)), "") {
		var f File
		size, _ := strconv.Atoi(n)
		f.index = index
		f.size = size
		index += size
		if i % 2 == 0 {
			f.id = i/2 
			files = append(files, f)
			for j := 0; j < size; j++ {
				final = append(final, i / 2)
			}
			b += size
		} else {
			f.id = -1
			spaces = append(spaces, f)
			for j := 0; j < size; j++ {
				final = append(final, -1)
			}
		}
	}


	if part2 {
		fmt.Println(final)
		for j := len(files) - 1; j > 0; j-- {
			f := &files[j]
			for k := 0; k < len(spaces) - 1; k++ {
				s := &spaces[k]
				if f.size <= s.size && f.index > s.index {
					for f.size > 0 {
						final[s.index] = f.id
						final[f.index] = s.id
						f.size--
						f.index++
						s.size--
						s.index++
					}
					break
				}
			}
		}

		for c, d := range final {
			if d != -1 {
				sum += c * d
			}
		}

	} else {

		final = final[:b]
		a = len(files) - 1

		for files[a].index + files[a].size + 1> spaces[0].index {
			final[spaces[0].index] = files[a].id

			if spaces[0].size == 1 {
				spaces = spaces[1:]
			} else {
				spaces[0].size--
				spaces[0].index++
			}

			if files[a].size == 1 {
				files = files[:a]
				a-- 
			} else {
				files[a].size--
			}
		}	

		for c, d := range final {
			sum += d * c	
		}  
	}

	return sum
}
