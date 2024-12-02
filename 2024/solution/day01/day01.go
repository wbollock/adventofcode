package main

import (
	"adventofcode/2024/utils"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

func main() {

	input, err := utils.LoadInput("day01")
	if err != nil {
		log.Print(err)
	}
	result := processInputPart1(input)
	fmt.Println(result)
}

func processInputPart1(input string) (result int) {

	inputSlice := strings.Split(string(input), "\n")
	var numberSlice []int
	var numberSlice2 []int

	for _, line := range inputSlice {
		numbersSlice := strings.Split(line, " ")

		num1, err := strconv.Atoi(string(numbersSlice[0]))
		if err != nil {
			log.Print(err)
		}
		numberSlice = append(numberSlice, num1)

		num2, err := strconv.Atoi(string(numbersSlice[3]))
		if err != nil {
			log.Print(err)
		}

		numberSlice2 = append(numberSlice2, num2)
		log.Print(numbersSlice)
	}

	sort.Ints(numberSlice)
	sort.Ints(numberSlice2)

	var distanceSlice []int

	for _, num := range numberSlice {
		var foundCounter int
		for _, lameNum := range numberSlice2 {
			if num == lameNum {
				foundCounter++
			}
		}
		distance := num * foundCounter
		// distance := num - numberSlice2[i]
		// if distance < 0 {
		// 	distance = distance * -1
		// }
		// log.Print(distance)
		distanceSlice = append(distanceSlice, distance)
	}

	for _, num := range distanceSlice {
		result += num
	}

	return result
}
