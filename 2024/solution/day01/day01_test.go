package main

import (
	"adventofcode/2024/utils"
	"log"
	"os"
	"testing"
)

func TestPart1(t *testing.T) {
	// expected := 11
	expected := 31

	// lazy soluton to deal with test cwd
	err := os.Chdir("../../../")
	if err != nil {
		log.Fatalf("Failed to change directory: %v", err)
	}
	input, err := utils.LoadInput("day01test")
	if err != nil {
		log.Print(err)
	}
	result := processInputPart1(input)

	if result != expected {
		t.Errorf("got %d; want %d", result, expected)
	}
}
