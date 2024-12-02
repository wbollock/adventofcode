package main

import (
	"adventofcode/2024/utils"
	"log"
	"strconv"
	"strings"
)

func main() {

	input, err := utils.LoadInput("day02")
	if err != nil {
		log.Print(err)
	}
	result := processInputPart1(input)
	log.Print("Part 1 ", result)
	result = processInputPart2(input)
	log.Print("Part 2 ", result)
}

func isLevelSequenceValidPart1(levelSlice []int) bool {
	var levelOkay bool
	increasing, decreasing := false, false

	for i, num := range levelSlice {
		if i+1 >= len(levelSlice) {
			break
		}

		difference := num - levelSlice[i+1]
		if difference == 0 {
			return false
		}

		if difference < 0 {
			difference = -difference
			decreasing = true
		} else {
			increasing = true
		}

		if increasing && decreasing {
			return false
		}

		if difference >= 1 && difference <= 3 {
			levelOkay = true
			log.Print("level okay ", difference, num, levelSlice[i+1])
		} else {
			log.Print("level not okay ", difference, num, levelSlice[i+1])
			return false
		}
	}

	return levelOkay
}

// func isLevelSequenceValidPart2(levelSlice []int) bool {

// 	increasing, decreasing := false, false

// 	strikes := 1

// 	for i, num := range levelSlice {
// 		if i+1 >= len(levelSlice) {
// 			break
// 		}

// 		difference := num - levelSlice[i+1]
// 		if difference == 0 {
// 			strikes--
// 			break
// 		}

// 		if difference < 0 {
// 			difference = -difference
// 			decreasing = true
// 		} else {
// 			increasing = true
// 		}

// 		if increasing && decreasing {
// 			strikes--
// 			break
// 		}

// 		if difference >= 1 && difference <= 3 {
// 			log.Print("level okay ", difference, num, levelSlice[i+1])
// 		} else {
// 			log.Print("level not okay ", difference, num, levelSlice[i+1])
// 			strikes--
// 			break
// 		}
// 	}
// 	if strikes == 0 {
// 		return false
// 	} else {
// 		return true
// 	}
// }

func processInputPart1(input string) (result int) {

	inputSlice := strings.Split(string(input), "\n")

	var safeCount int
	// reports are row, level is each column
	for _, report := range inputSlice {
		numbersSlice := strings.Split(report, " ")
		log.Print(numbersSlice)

		var levelSlice []int

		for _, level := range numbersSlice {
			if num, err := strconv.Atoi(level); err == nil {
				levelSlice = append(levelSlice, num)
			}
		}

		// 1. must be all increasing or decreasing
		// 2. Any two adjacent levels differ by at least one and at most three
		levelOkay := isLevelSequenceValidPart1(levelSlice)
		if levelOkay {
			safeCount++
		}
	}
	return safeCount
}

func processInputPart2(input string) (result int) {

	inputSlice := strings.Split(string(input), "\n")

	var safeCount int
	// reports are row, level is each column
	for _, report := range inputSlice {
		numbersSlice := strings.Split(report, " ")
		log.Print(numbersSlice)

		var levelSlice []int

		for _, level := range numbersSlice {
			if num, err := strconv.Atoi(level); err == nil {
				levelSlice = append(levelSlice, num)
			}
		}

		// 1. must be all increasing or decreasing
		// 2. Any two adjacent levels differ by at least one and at most three
		levelOkay := isLevelSequenceValidPart1(levelSlice)
		if levelOkay {
			safeCount++
		} else {
			// brute force removing a random element if level is not valid
			lenSlice := len(levelSlice)
			for i := range lenSlice {
				// https://go.dev/wiki/SliceTricks#filtering-without-allocating
				// make sure to create new slice with size pre-allocated
				levelSliceRemoved := make([]int, len(levelSlice))
				copy(levelSliceRemoved, levelSlice)

				levelSliceRemoved = append(levelSliceRemoved[:i], levelSliceRemoved[i+1:]...)

				levelOkay = isLevelSequenceValidPart1(levelSliceRemoved)
				if levelOkay {
					log.Print("level okay with element removed ", levelSliceRemoved, " original level", levelSlice)
					safeCount++
					break
				} else {
					log.Print("level not okay with element removed ", levelSliceRemoved, " original level", levelSlice)
				}
			}
		}
	}
	return safeCount
}
