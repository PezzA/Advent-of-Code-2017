package day02

import "testing"

func Test_Part1(t *testing.T) {
	key := PartOne(day2PuzzleInput())
	if key != "74921" {
		t.Error("Expected 74921 got ", key)
	}
}

func Test_Part2(t *testing.T) {
	key := PartTwo(day2PuzzleInput())
	if key != "A6B35" {
		t.Error("Expected A6B35 got ", key)
	}
}
