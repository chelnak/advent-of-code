package main

import (
	"advent-of-code/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {

	var lines []string
	var horizontal int
	var depth int
	var aim int
	var result int

	utils.Read("inputs.txt", &lines)

	// Part 1
	for _, line := range lines {

		slice := strings.Fields(line)
		command := slice[0]
		position, _ := strconv.Atoi(slice[1])

		switch command {
		case "forward":
			horizontal = horizontal + position
		case "down":
			depth = depth + position
		case "up":
			depth = depth - position
		}

	}

	result = horizontal * depth
	fmt.Println("Part 1:")
	fmt.Println(result)

	// Part 2
	horizontal = 0
	depth = 0
	result = 0

	for _, line := range lines {

		slice := strings.Fields(line)
		command := slice[0]
		position, _ := strconv.Atoi(slice[1])

		switch command {
		case "forward":
			horizontal = horizontal + position
			depth = depth + (aim * position)
		case "down":
			aim = aim + position
		case "up":
			aim = aim - position
		}

	}

	result = horizontal * depth
	fmt.Println("Part 2:")
	fmt.Println(result)
}
