package day02

import "testing"

func Test_Part1(t *testing.T) {
	result, err := PartOne()

	if err != nil {
		t.Fatal("Failed testing part one", err)
	}

	if result != "74921" {
		t.Error("Expected 74921 got ", result)
	}
}

func Test_Part2(t *testing.T) {
	result, err := PartTwo()

	if err != nil {
		t.Fatal("Failed testing part two", err)
	}

	if result != "A6B35" {
		t.Error("Expected A6B35 got ", result)
	}
}
