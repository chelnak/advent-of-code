package main

import (
	"advent-of-code/utils"
	"fmt"
)

func simulateFishPopulationGrowth(data []int, maxDays int) int {

	maxDaysInCycle := 9

	fish := make([]int, maxDaysInCycle)

	for _, f := range data {
		fish[f]++
	}

	for i := 0; i < maxDays; i++ {
		count := make([]int, maxDaysInCycle)
		for i := 1; i < maxDaysInCycle; i++ {
			count[i-1] = fish[i]
		}
		count[6] += fish[0]
		count[8] += fish[0]
		fish = count
	}

	numberOfFish := utils.Sum(fish)

	fmt.Println("Number of fish after", maxDays, "days:", numberOfFish)
	return numberOfFish
}

func main() {
	lines := []string{}
	utils.Read("inputs.txt", &lines)

	fish := utils.SplitStringToInt(lines[0], ",")

	simulateFishPopulationGrowth(fish, 18)
	simulateFishPopulationGrowth(fish, 80)
	simulateFishPopulationGrowth(fish, 256)

}
