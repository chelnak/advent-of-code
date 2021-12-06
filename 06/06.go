package main

import (
	"advent-of-code/utils"
	"fmt"
)

func sumulateFishPopulationGrowth(fish []int, maxDays int) {

	currentDay := 1

	fishMap := make(map[int]int)

	for idx, f := range fish {
		fishMap[idx] = f
	}

	for currentDay <= maxDays {

		idx := 0
		currentLength := len(fishMap)

		for idx <= currentLength {

			if _, ok := fishMap[idx]; ok {

				if fishMap[idx] == 0 {

					fishMap[idx] = 6

					fishMap[len(fishMap)+1] = 8

				} else {

					fishMap[idx]--
				}
			}

			idx++

		}

		currentDay++
	}

	numberOfFish := len(fishMap)

	fmt.Println("Number of fish after", maxDays, "days:", numberOfFish)
}

func main() {
	lines := []string{}
	utils.Read("inputs2.txt", &lines)

	fish := utils.SplitStringToInt(lines[0], ",")

	sumulateFishPopulationGrowth(fish, 18)
	sumulateFishPopulationGrowth(fish, 256)

}
