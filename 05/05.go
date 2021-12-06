package main

import (
	"advent-of-code/utils"
	"fmt"
	"math"
)

type Point struct {
	x1, y1 int
	x2, y2 int
}

func main() {

	lines := []string{}
	vhlines := []Point{}

	utils.Read("inputs.txt", &lines)

	var points = make([]Point, len(lines))

	maxX := 0
	maxY := 0

	for idx, line := range lines {
		fmt.Sscanf(line, "%d,%d -> %d,%d",
			&points[idx].x1,
			&points[idx].y1,
			&points[idx].x2,
			&points[idx].y2)

		if points[idx].x1 > maxX {
			maxX = points[idx].x1
		}

		if points[idx].x2 > maxX {
			maxX = points[idx].x2
		}

		if points[idx].y1 > maxY {
			maxY = points[idx].y1
		}

		if points[idx].y2 > maxY {
			maxY = points[idx].y2
		}

		if points[idx].x1 == points[idx].x2 || points[idx].y1 == points[idx].y2 {
			vhlines = append(vhlines, points[idx])
		}

	}

	grid := make([][]int, 1000)
	for i := 0; i < 1000; i++ {
		grid[i] = make([]int, 1000)
	}

	for _, line := range points {

		if line.x1 == line.x2 {

			y1 := line.y1
			y2 := line.y2

			if y2 < y1 {
				y1 = line.y2
				y2 = line.y1
			}

			for y := y1; y <= y2; y++ {
				grid[y][line.x1]++
			}

		} else if line.y1 == line.y2 {

			x1 := line.x1
			x2 := line.x2
			if x2 < x1 {
				x1 = line.x2
				x2 = line.x1
			}

			for x := x1; x <= x2; x++ {
				grid[line.y1][x]++
			}

		} else {

			dx := int(float64(line.x2-line.x1) / math.Abs(float64(line.x2-line.x1)))
			dy := int(float64(line.y2-line.y1) / math.Abs(float64(line.y2-line.y1)))

			y := line.y1
			x := line.x1

			for y != line.y2+dy {
				grid[y][x]++
				y += dy
				x += dx
			}
		}

	}

	count := 0
	for _, row := range grid {

		for _, v := range row {
			if v >= 2 {
				count++
			}
		}
	}

	fmt.Println(count)

}
