package main

import (
	"adventofcode/2024/utils"
	"fmt"
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
	fmt.Println(result)
}

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
		var levelOkay bool
		increasing, decreasing := false, false
		for i, num := range levelSlice {
			if i+1 >= len(levelSlice) {
				break
			}
			if num-levelSlice[i+1] == 0 {
				levelOkay = false
				break
			}

			difference := num - levelSlice[i+1]
			if difference < 0 {
				difference = difference * -1
				decreasing = true
			} else {
				increasing = true
			}
			if increasing && decreasing {
				levelOkay = false
				break
			}
			if difference >= 1 && difference <= 3 {
				levelOkay = true
				log.Print("level okay ", difference, num, levelSlice[i+1])
			} else {
				levelOkay = false
				log.Print("level not okay ", difference, num, levelSlice[i+1])
				break
			}
		}
		if levelOkay {
			safeCount++
		}
	}
	return safeCount
}
