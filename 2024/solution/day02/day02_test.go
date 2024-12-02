package main

import (
	"adventofcode/2024/utils"
	"log"
	"os"
	"testing"
)

func TestPart1(t *testing.T) {
	expected := 2

	err := os.Chdir("../../../")
	if err != nil {
		log.Fatalf("Failed to change directory: %v", err)
	}
	input, err := utils.LoadInput("day02test")
	if err != nil {
		log.Print(err)
	}
	result := processInputPart1(input)

	if result != expected {
		t.Errorf("got %d; want %d", result, expected)
	}
}
