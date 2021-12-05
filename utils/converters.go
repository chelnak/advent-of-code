package utils

import (
	"strconv"
	"strings"
)

func Convert(input string) (int64, error) {
	return strconv.ParseInt(input, 2, 64)
}

func SplitStringToInt(input string, sep string) []int {

	var result []int
	var chars []string

	if sep == "" {
		chars = strings.Fields(input)
	} else {
		chars = strings.Split(input, sep)
	}

	for _, char := range chars {
		val, _ := strconv.Atoi(char)
		result = append(result, val)
	}

	return result

}
