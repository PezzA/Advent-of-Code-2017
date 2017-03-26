package day01

import "testing"

func Test_Part1(t *testing.T) {
	result, err := PartOne(PuzzleInput())

	if err != nil {
		t.Fatal("Failed testing part one", err)
	}

	if result != "241" {
		t.Error("Expected 241 got ", result)
	}
}

func Test_Part2(t *testing.T) {
	result, err := PartTwo(PuzzleInput())

	if err != nil {
		t.Fatal("Failed testing part two", err)
	}

	if result != "116" {
		t.Error("Expected 116 got ", result)
	}
}
