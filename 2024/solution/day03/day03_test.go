package main

import (
	"adventofcode/2024/utils"
	"log"
	"os"
	"testing"
)

func TestPart1(t *testing.T) {
	expected := 161

	err := os.Chdir("../../../")
	if err != nil {
		log.Fatalf("Failed to change directory: %v", err)
	}
	input, err := utils.LoadInput("day03test")
	if err != nil {
		log.Print(err)
	}
	result := processInputPart1(input)

	if result != expected {
		t.Errorf("got %d; want %d", result, expected)
	}
}

func TestPart2(t *testing.T) {
	expected := 48

	// err := os.Chdir("../../../")
	// if err != nil {
	// 	log.Fatalf("Failed to change directory: %v", err)
	// }
	input, err := utils.LoadInput("day03test")
	if err != nil {
		log.Print(err)
	}
	result := processInputPart2(input)

	if result != expected {
		t.Errorf("got %d; want %d", result, expected)
	}
}

// func TestPart2Special(t *testing.T) {
// 	expected := 3971651

// 	// err := os.Chdir("../../../")
// 	// if err != nil {
// 	// 	log.Fatalf("Failed to change directory: %v", err)
// 	// }
// 	input, err := utils.LoadInput("day03test3")
// 	if err != nil {
// 		log.Print(err)
// 	}
// 	result := processInputPart2(input)

// 	if result != expected {
// 		t.Errorf("got %d; want %d", result, expected)
// 	}
// }
