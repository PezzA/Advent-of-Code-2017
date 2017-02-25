package main

import (
	"fmt"

	"github.com/pezza/advent-of-go/day01"
	"github.com/pezza/advent-of-go/day02"
	"github.com/pezza/advent-of-go/day03"
	"github.com/pezza/advent-of-go/day04"

	"os"
	"strconv"
)

type puzzle func() (string, error)

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
		runDay(day01.PartOne, day01.PartTwo)
	case 2:
		runDay(day02.PartOne, day02.PartTwo)
	case 3:
		runDay(day03.PartOne, day03.PartTwo)
	case 4:
		runDay(day04.PartOne, day04.PartTwo)
	}
}

func runDay(partOne puzzle, partTwo puzzle) {
	partOneResult, partOneErr := partOne()
	fmt.Print("Part One ")
	if partOneErr != nil {
		fmt.Print(partOneErr)
	} else {
		fmt.Print(partOneResult)
	}
	fmt.Print("\n")
	partTwoResult, partTwoErr := partTwo()

	fmt.Print("Part Two ")
	if partTwoErr != nil {
		fmt.Print(partTwoErr)
	} else {
		fmt.Print(partTwoResult)
	}
	fmt.Print("\n")
}
