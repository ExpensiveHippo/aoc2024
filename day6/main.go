package main

import (
	"os"
	"fmt"
	"image"
	"strings"
	"slices"
)

type Guard struct {
	pos image.Point;
	dir int;
}

func main() {
	directions := "^>v<"
	moves := []image.Point{
		{ 0, -1 },
		{ 1, 0 },
		{ 0, 1 }, 
		{ -1, 0 },
	}
	m := map[string]int {
		"#": -1,
		".": 0, 
	}
	file, _ := os.ReadFile("day6/input.txt")
	var guard Guard
	grid := map[image.Point]int{}
	for y, line := range strings.Fields(string(file)) {
		for x, o := range strings.Split(line, "") {
			pt := image.Pt(x, y)
			if dir := strings.IndexAny(o, directions); dir != -1 {
				guard.pos = pt
				guard.dir = dir
				grid[pt] = 1
			} else {
				grid[pt] = m[o]
			}
		}
	}
	
	simulate := func(guard Guard, newObs image.Point) map[image.Point][]int {
		seen := map[image.Point][]int{}	

		for {
			if _, inBounds := grid[guard.pos]; !inBounds {
				return seen
			}

			// stuck in loop	
			if slices.Contains(seen[guard.pos], guard.dir) {
				return nil
			}
			
			seen[guard.pos] = append(seen[guard.pos], guard.dir)
			np := guard.pos.Add(moves[guard.dir])

			if grid[np] == -1 || np == newObs {
				guard.dir = (guard.dir + 1) % len(directions)
			} else {
				guard.pos = np
			}
		}
	} 

	steps, count := simulate(guard, image.Pt(-1, -1)), 0
	
	for pos, _ := range steps {
		if pos != guard.pos && simulate(guard, pos) == nil {
			count++
		}
	}

	fmt.Println("Part 1: ", len(steps))
	fmt.Println("Part 2: ", count)
}

