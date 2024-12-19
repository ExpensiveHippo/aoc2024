package main

import (
	"os"
	"fmt"
	"strings"
	"image"
	"math"
)


func main() {
	file, _ := os.ReadFile("day8/input.txt")
	antennas := map[string][]image.Point{}
	antinodes1 := map[image.Point]int{}
	antinodes2 := map[image.Point]int{}

	s := strings.Fields(string(file))
	bound := image.Pt(len(s[0]), len(s))

	for y, line := range s {
		for x, node := range strings.Split(line, "") {
			if node != "." {
				antennas[node] = append(antennas[node], image.Pt(x, y))
			}
		}
	}

	for _, locations := range antennas {
		nodes1 := findAntinodes1(locations, bound)
		nodes2 := findAntinodes2(locations, bound)
		for node, _ := range nodes1 {
			antinodes1[node] = 1
		}
		for node, _ := range nodes2 {
			antinodes2[node] = 1
		}
	}

	fmt.Println("Part 1: ", len(antinodes1))
	fmt.Println("Part 2: ", len(antinodes2) )
}




func findAntinodes1(locations []image.Point, bound image.Point) map[image.Point]int {

	if len(locations) < 2 {
		return nil
	}
	
	curr := locations[0]
	antinodes := map[image.Point]int{}

	outOfBounds := func(pt image.Point) bool {
		return (pt.X < 0 ||
		pt.Y < 0 ||
		pt.X > bound.X - 1 ||
		pt.Y > bound.Y - 1)
	}
	
	for _, loc := range locations[1:] {
		xDis := curr.X - loc.X
		yDis := curr.Y - loc.Y

		// check between both points
		if xDis % 3 == 0 && yDis % 3 == 0 {
			xInt := xDis / 3
			yInt := yDis / 3
			antinodes[image.Pt(curr.X - xInt, curr.Y - yInt)] = 1	
			antinodes[image.Pt(curr.X - xInt * 2, curr.Y - yInt * 2)] = 1
		} 
		
		// if 2d is the distance between the points,
		// check for 2d away the points
		anti1 := image.Pt(curr.X + xDis, curr.Y + yDis) 	
		anti2 := image.Pt(loc.X - xDis, loc.Y - yDis)	

		if !outOfBounds(anti1) {
			antinodes[anti1] = 1
		}

		if !outOfBounds(anti2) {
			antinodes[anti2] = 1
		}
	}

	if a := findAntinodes1(locations[1:], bound); a != nil {
		for pt, _ := range a {
			antinodes[pt] = 1 
		}	
	}

	return antinodes
}

func findAntinodes2(locations []image.Point, bound image.Point) map[image.Point]int {

	if len(locations) < 2 {
		return nil
	}
	
	curr := locations[0]
	antinodes := map[image.Point]int{ curr: 1 }

	outOfBounds := func(pt image.Point) bool {
		return (pt.X < 0 ||
		pt.Y < 0 ||
		pt.X > bound.X - 1 ||
		pt.Y > bound.Y - 1)
	}
	
	for _, loc := range locations[1:] {
		xDis := curr.X - loc.X
		yDis := curr.Y - loc.Y
		m := int(math.Min(math.Abs(float64(xDis)), math.Abs(float64(yDis))))
		var interval image.Point

		if xDis % m == 0 && yDis % m == 0 {
			interval = image.Pt(xDis/m, yDis/m)
		} else {
			interval = image.Pt(xDis, yDis)
		}

		node := curr

		for {
			node = node.Add(interval)

			if outOfBounds(node) {
				break
			}

			antinodes[node] = 1
		}

		node = curr

		for {
			node = node.Sub(interval)

			if outOfBounds(node) {
				break
			}

			antinodes[node] = 1
		}
	}
	

	if a := findAntinodes2(locations[1:], bound); a != nil {
		for pt, _ := range a {
			antinodes[pt] = 1 
		}	
	}

	return antinodes
}


