package main

import (
	"adventofcode/2024/utils"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	input, err := utils.LoadInput("day03")
	if err != nil {
		log.Print(err)
	}
	// result := processInputPart1(input)
	// log.Print("Part 1 ", result)
	result := processInputPart2(input)
	log.Print("Part 2 ", result)
}

type Memory struct {
	Mults []string
}

func gatherMults(input string) (mults []string) {
	// find all valid mults in a string
	// regex boi go brrr
	re := regexp.MustCompile("mul\\(\\d+,\\d+\\)")
	mults = re.FindAllString(input, -1)
	log.Print("regex match ", mults)
	return mults
}

func gatherMultsPart2(input string) []string {

    // now remove all text between a don't() and do()

	re := regexp.MustCompile("don't\\(\\)(.*?)do\\(\\)")
	log.Print("input BEFORE ", input)
	goodMults := re.ReplaceAllString(input, "")
	// then remove everything from don't to end of line
	re = regexp.MustCompile("don't\\(\\).*")
	goodMults = re.ReplaceAllString(goodMults, "")
	log.Print("INPUT AFTER ", goodMults)

	// get mults as normal
	re = regexp.MustCompile("mul\\(\\d+,\\d+\\)")

	mults := re.FindAllString(goodMults, -1)

	// goodMults := re.ReplaceAllString(mults, "")
	log.Print("mults are ", mults)
	return mults
}

func (m *Memory) multiplyMults() int {
	// do the actual mult calcs
	var result int

	for _, mults := range m.Mults {
		multSlice := strings.Split(string(mults), ",")
		var digitSlice []int
		for _, mult := range multSlice {
			// log.Print("mult ", mult)
			re := regexp.MustCompile("\\d+")
			digits := re.FindAllString(mult, -1)
			for _, digit := range digits {
				num, err := strconv.Atoi(digit)
				if err != nil {
					log.Print(err)
				}
				digitSlice = append(digitSlice, num)
			}
		}
		multResult := 1
		log.Print("digitSlice ", digitSlice)
		for _, digit := range digitSlice {
			multResult *= digit
			// log.Print("mulitplying ", digit)
		}
		result += multResult
		log.Print("adding result ", result, " multresult ", multResult)
	}

	return result
}

func processInputPart1(input string) (result int) {

	inputSlice := strings.Split(string(input), "\n")

	var mults Memory

	for _, row := range inputSlice {

		mults.Mults = append(mults.Mults, gatherMults(row)...)
	}
	result = mults.multiplyMults()
	return result
}

func processInputPart2(input string) (result int) {

	inputSlice := strings.Split(string(input), "\n")

	var mults Memory

	for _, row := range inputSlice {
		mults.Mults = append(mults.Mults, gatherMultsPart2(row)...)
	}
	result = mults.multiplyMults()
	return result
}
