package main

import (
	"fmt"

	"github.com/pezza/advent-of-go/day01"
	"github.com/pezza/advent-of-go/day02"
	"github.com/pezza/advent-of-go/day03"
	"github.com/pezza/advent-of-go/day04"
	"github.com/pezza/advent-of-go/day05"
	"github.com/pezza/advent-of-go/day06"
	"github.com/pezza/advent-of-go/day07"
	"github.com/pezza/advent-of-go/day09"
	"github.com/pezza/advent-of-go/day14"
	"github.com/pezza/advent-of-go/day15"

	"os"
	"strconv"
)

type puzzle func(string) (string, error)

type dayRunner interface {
	PartOne() (string, error)
	PartTwo() (string, error)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("USEAGE: advent-of-go <day>")
		fmt.Println("<day> = day number of puzzle to run.  e.g. advent-of-go 2")

		//TODO - need to ensure value is actuall a string
		return
	}

	switch day, _ := strconv.Atoi(os.Args[1]); day {
	case 1:
		runDay(day01.PartOne, day01.PartTwo, day01.PuzzleInput())
	case 2:
		runDay(day02.PartOne, day02.PartTwo, day02.PuzzleInput())
	case 3:
		runDay(day03.PartOne, day03.PartTwo, day03.PuzzleInput())
	case 4:
		runDay(day04.PartOne, day04.PartTwo, day04.PuzzleInput())
	case 5:
		runDay(day05.PartOne, day05.PartTwo, day05.PuzzleInput())
	case 6:
		runDay(day06.PartOne, day06.PartTwo, day06.PuzzleInput())
	case 7:
		runDay(day07.PartOne, day07.PartTwo, day07.PuzzleInput())
	case 9:
		runDay(day09.PartOne, day09.PartTwo, day09.PuzzleInput())
	case 14:
		runDay(day14.PartOne, day14.PartTwo, day14.PuzzleInput())
	case 15:
		runDay(day15.PartOne, day15.PartTwo, day15.PuzzleInput())
	}
}

func runDay(partOne puzzle, partTwo puzzle, input string) {
	partOneResult, partOneErr := partOne(input)
	fmt.Print("Part One ")
	if partOneErr != nil {
		fmt.Print(partOneErr)
	} else {
		fmt.Print(partOneResult)
	}
	fmt.Print("\n")
	partTwoResult, partTwoErr := partTwo(input)

	fmt.Print("Part Two ")
	if partTwoErr != nil {
		fmt.Print(partTwoErr)
	} else {
		fmt.Print(partTwoResult)
	}
	fmt.Print("\n")
}
