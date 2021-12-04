package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readFile(file string, lines *[]string) {

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
		value := scanner.Text()
		*lines = append(*lines, value)
	}
}

func getMostCommonBitAtPos(lines []string, pos int, invert bool) int {
	high := 1
	low := 0

	count := make(map[int]int)

	if invert {
		high = 0
		low = 1
	}

	for _, line := range lines {
		bits := strings.Split(line, "")
		key, _ := strconv.Atoi(bits[pos])
		count[key]++
	}

	if count[0] == count[1] {
		return high
	} else if count[0] > count[1] {
		return low
	} else {
		return high
	}
}

func getRating(lines []string, width int, invert bool) []string {

	pos := 0

	for pos < width {

		if len(lines) == 1 {
			break
		}

		var common []string
		for _, number := range lines {

			bits := strings.Split(number, "")
			bit, _ := strconv.Atoi(bits[pos])

			commonBitForPos := getMostCommonBitAtPos(lines, pos, invert)

			if bit == commonBitForPos {
				common = append(common, number)
			}
		}

		lines = common
		pos++

	}

	return lines

}

func main() {

	var lines []string
	var gamma []string
	var epsilon []string

	readFile("inputs.txt", &lines)

	pos := 0
	width := len(lines[0])

	// Part 1
	for pos < width {

		bit := getMostCommonBitAtPos(lines, pos, false)

		if bit == 0 {
			gamma = append(gamma, "0")
			epsilon = append(epsilon, "1")
		} else {
			gamma = append(gamma, "1")
			epsilon = append(epsilon, "0")
		}
		pos++
	}

	gamma_decimal, _ := strconv.ParseInt(strings.Join(gamma, ""), 2, 64)
	epsilon_decimal, _ := strconv.ParseInt(strings.Join(epsilon, ""), 2, 64)

	fmt.Println(gamma_decimal * epsilon_decimal)

	// Part 2
	co2ScrubberRating_decimal, _ := strconv.ParseInt(strings.Join(getRating(lines, width, false), ""), 2, 64)
	oxygenGeneratorRating_decimal, _ := strconv.ParseInt(strings.Join(getRating(lines, width, true), ""), 2, 64)
	fmt.Println(oxygenGeneratorRating_decimal * co2ScrubberRating_decimal)
}
