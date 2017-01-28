package day01

import "testing"

// BenchmarkPartOne tests the performance of the first day part
func BenchmarkPartOne(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		PartOne()
	}
}

func Test_Part1(t *testing.T) {
	result, err := PartOne()

	if err != nil {
		t.Fatal("Failed testing part one", err)
	}

	if result != "241" {
		t.Error("Expected 241 got ", result)
	}
}

func Test_Part2(t *testing.T) {
	result, err := PartTwo()

	if err != nil {
		t.Fatal("Failed testing part two", err)
	}

	if result != "116" {
		t.Error("Expected 116 got ", result)
	}
}
