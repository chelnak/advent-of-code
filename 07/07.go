package main

import (
	"advent-of-code/utils"
	"fmt"
	"math"
)

func getPart2ShortestDistance(positions []int, max int) int {
	var distances []int = make([]int, max+1)
	for i := 0; i <= max; i++ {

		for _, value := range positions {
			base := int(math.Abs(float64(value - i)))
			triangular := (base * (base + 1)) / 2
			distances[i] += triangular
		}
	}

	return getMinOf(distances)
}

func getPart1ShortestDistance(positions []int, max int) int {

	var distances []int = make([]int, max+1)
	for i := 0; i <= max; i++ {
		for _, value := range positions {
			distances[i] += int(math.Abs(float64(value - i)))
		}
	}

	return getMinOf(distances)
}

func getMaxOf(values []int) int {

	max := 0
	for _, value := range values {

		if value > max {
			max = value
		}
	}

	return max
}

func getMinOf(values []int) int {

	min := math.MaxInt32
	for _, value := range values {
		if value < min {
			min = value
		}
	}

	return min
}

func main() {

	lines := []string{}
	utils.Read("inputs.txt", &lines)

	positions := utils.SplitStringToInt(lines[0], ",")
	max := getMaxOf(positions)

	fmt.Println("Part 1:", getPart1ShortestDistance(positions, max))
	fmt.Println("Part 2:", getPart2ShortestDistance(positions, max))

}
