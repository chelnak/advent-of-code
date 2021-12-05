package main

import (
	"advent-of-code/utils"
	"fmt"
	"strconv"
	"strings"
)

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

func getLifeSupportRating(lines []string, width int, invert bool) string {

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

	return strings.Join(lines, "")
}

func getPowerRating(lines []string, width int, invert bool) string {

	pos := 0
	var result []string
	for pos < width {

		bit := getMostCommonBitAtPos(lines, pos, invert)

		if bit == 0 {
			result = append(result, "0")
		} else {
			result = append(result, "1")
		}
		pos++

	}

	return strings.Join(result, "")
}

func main() {
	var lines []string
	utils.Read("inputs.txt", &lines)
	width := len(lines[0])

	// Part 1
	gammaRate, _ := utils.Convert(getPowerRating(lines, width, false))
	epsilonRate, _ := utils.Convert(getPowerRating(lines, width, true))
	fmt.Println("Power rating:", gammaRate*epsilonRate)

	// Part 2
	co2ScrubberRating, _ := utils.Convert(getLifeSupportRating(lines, width, false))
	oxygenGeneratorRating, _ := utils.Convert(getLifeSupportRating(lines, width, true))
	fmt.Println("Life support rating:", co2ScrubberRating*oxygenGeneratorRating)
}
