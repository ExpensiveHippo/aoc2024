package main 

import (
	"os"
	"fmt"
	"image"
	"strings"
	"strconv"
)

func main() {
	file, _ := os.ReadFile("day10/input.txt")
	grid := map[image.Point]int{}
	trailHeads := []image.Point{}

	for y, line := range strings.Fields(string(file)) {
		line = strings.TrimSpace(line)
		for x, level := range strings.Split(line, ""){
			l, _ := strconv.Atoi(level)
			grid[image.Pt(x, y)] = l
			if l == 0 {
				trailHeads = append(trailHeads, image.Pt(x, y))
			}
		}
	}
	
	count, count2 := 0, 0
	
	for _, pos := range trailHeads {
		count += countPath(pos, grid)
		count2 += countRating(pos, grid)
	}
	fmt.Println("Part 1: ", count)
	fmt.Println("Part 2: ", count2)
}


func countPath(pos image.Point, grid map[image.Point]int) int {
	seen := map[image.Point]int{}
	return helper(0, pos, seen, grid)	
}

func helper(level int, pos image.Point, seen, grid map[image.Point]int) int {
	if level == 9 {
		seen[pos] = 1
		return 1
	}	

	moves := []image.Point {
		{ 0, -1 },
		{ 1, 0 },
		{ 0, 1 },
		{ -1, 0 },
	}
	sum := 0

	for _, dir := range moves {
		if next := pos.Add(dir); grid[next] == level+1 && seen[next] != 1 {
			sum += helper(level+1, next, seen, grid)
		}
	}

	return sum

}

func countRating(pos image.Point, grid map[image.Point]int) int {
	return helper2(0, pos, grid)
}

func helper2(level int, pos image.Point, grid map[image.Point]int) int {
	if level == 9 {
		return 1
	}	

	moves := []image.Point {
		{ 0, -1 },
		{ 1, 0 },
		{ 0, 1 },
		{ -1, 0 },
	}
	sum := 0

	for _, dir := range moves {
		if next := pos.Add(dir); grid[next] == level+1 {
			sum += helper2(level+1, next, grid)
		}
	}

	return sum

}

