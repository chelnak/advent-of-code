package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readFile(file string, lines *[]int) {

	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		value, _ := strconv.Atoi(scanner.Text())
		*lines = append(*lines, value)
	}
}

func sum(numbers []int) int {
	var result int
	for _, number := range numbers {
		result += number
	}
	return result
}

func makeWindow(lines []int, windowSize int) [][]int {
	var result [][]int
	for i := 0; i < len(lines)-windowSize+1; i++ {
		result = append(result, lines[i:i+windowSize])
	}
	return result
}

func main() {

	var lines []int
	var count int
	var last int
	var lastSum int
	var windows [][]int

	readFile("inputs.txt", &lines)

	// Part 1
	last = lines[0]
	for _, line := range lines {
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
	windows = makeWindow(lines, 3)
	lastSum, windows = sum(windows[0]), windows[1:]

	for _, group := range windows {
		currentSum := sum(group)

		if currentSum > lastSum {
			count++
		}

		lastSum = currentSum
	}

	fmt.Println("Part 2:")
	fmt.Println(count)

}
