package main

import (
	"advent-of-code/utils"
	"fmt"
	"strconv"
)

func makeWindow(lines []int, windowSize int) [][]int {
	var result [][]int
	for i := 0; i < len(lines)-windowSize+1; i++ {
		res := lines[i : i+windowSize]
		result = append(result, res)
	}
	return result
}

func main() {

	var lines []string
	var linesInt []int
	var count int
	var last int
	var lastSum int
	var windows [][]int

	utils.Read("inputs.txt", &lines)

	// convert to int
	for _, line := range lines {
		lineInt, _ := strconv.Atoi(line)
		linesInt = append(linesInt, lineInt)
	}

	// Part 1
	last = linesInt[0]
	for _, line := range linesInt {
		current := line
		if current > last {
			count++
		}
		last = current
	}

	fmt.Println("Part 1:")
	fmt.Println(count)

	// Part 2
	count = 0
	windows = makeWindow(linesInt, 3)
	lastSum, windows = utils.Sum(windows[0]), windows[1:]

	for _, group := range windows {
		currentSum := utils.Sum(group)

		if currentSum > lastSum {
			count++
		}

		lastSum = currentSum
	}

	fmt.Println("Part 2:")
	fmt.Println(count)

}
