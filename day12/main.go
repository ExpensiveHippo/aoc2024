package main

import (
	"os"
	"fmt"
	"strings"
	"image"
	"slices"
)

func main() {
	file, _ := os.ReadFile("day12/input.txt")
	grid := map[image.Point]string{}
	z := strings.Fields(strings.TrimSpace(string(file)))
	for y, line := range z {
		for x, c := range strings.Split(line, "") {
			grid[image.Pt(x, y)] = c
		}
	}

	moves := []image.Point {
		{ 0, -1 }, 
		{ 1, 0 }, 
		{ 0, 1 }, 
		{ -1, 0 }, 
	}

	outOfBounds := func(p image.Point) bool {
		return ( p.X < 0 ||
			p.Y < 0 ||
			p.X > len(z[0]) - 1 || 
			p.Y > len(z) - 1)
	}

	seen := map[image.Point]int{}
	sum1, sum2 := 0, 0

	for pt, char := range grid {
		if seen[pt] == 1 {
			continue
		}

		region := []image.Point { pt }
		sides := map[int][]image.Point{}
		area, peri := 0, 0

		for len(region) != 0 {
			curr := region[0]
			region = region[1:]
			if seen[curr] == 1 {
				continue
			}
			seen[curr] = 1
			area++

			for i, dir := range moves {
				newPt := curr.Add(dir)
				if outOfBounds(newPt) {
					peri++
					sides[i] = append(sides[i], curr)
					continue
				} 
				
				if seen[newPt] == 1 {
					if grid[newPt] != char {
						peri++
						sides[i] = append(sides[i], curr)
					}
				} else if grid[newPt] == char {
					region = append(region, newPt)
				} else {
					peri++
					sides[i] = append(sides[i], curr)
				}
			}
		}
		
		nSides := 0
		for i, edges := range sides {
			temp := map[int][]int{}
			if i % 2 == 0 {
				for _, edge := range edges {
					temp[edge.Y] = append(temp[edge.Y], edge.X) 	
				}
			} else {
				for _, edge := range edges {
					temp[edge.X] = append(temp[edge.X], edge.Y) 	
				}
			}

			// fmt.Println("Direction: ", i)
			// fmt.Println("Old: ", temp)

			for _, pos := range temp {
				slices.Sort(pos)	
				prev := pos[0]
				count := 1
				for len(pos) != 0 && count < len(pos){
					if pos[count] != prev + 1 {
						pos = pos[count:] 
						// fmt.Println("New: ", pos)
						nSides++
						count = 0
					}			
					prev = pos[count]
					count++
				}

				if len(pos) != 0 {
					nSides++
				}
			} 
		}
		sum1 += area * peri
		sum2 += area * nSides
		// fmt.Println(char)
		// fmt.Println("Area: ", area)
		// fmt.Println("Peri: ", peri)
		// fmt.Println("Sides: ", nSides)
	}
	fmt.Println("Part 1: ", sum1)
	fmt.Println("Part 2: ", sum2)
}


