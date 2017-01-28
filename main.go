package main

import (
	"fmt"

	"os"

	"github.com/pezza/advent-of-go/day01"
)

type dayRunner interface {
	PartOne(input string) (string, error)
	PartTwo(input string) (string, error)
}

func main() {
	var instructions = "R1, R1, R3, R1, R1, L2, R5, L2, R5, R1, R4, L2, R3, L3, R4, L5, R4, R4, R1, L5, L4, R5, R3, L1, R4, R3, L2, L1, R3, L4, R3, L2, R5, R190, R3, R5, L5, L1, R54, L3, L4, L1, R4, R1, R3, L1, L1, R2, L2, R2, R5, L3, R4, R76, L3, R4, R191, R5, R5, L5, L4, L5, L3, R1, R3, R2, L2, L2, L4, L5, L4, R5, R4, R4, R2, R3, R4, L3, L2, R5, R3, L2, L1, R2, L3, R2, L1, L1, R1, L3, R5, L5, L1, L2, R5, R3, L3, R3, R5, R2, R5, R5, L5, L5, R2, L3, L5, L2, L1, R2, R2, L2, R2, L3, L2, R3, L5, R4, L4, L5, R3, L4, R1, R3, R2, R4, L2, L3, R2, L5, R5, R4, L2, R4, L1, L3, L1, L3, R1, R2, R1, L5, R5, R3, L3, L3, L2, R4, R2, L5, L1, L1, L5, L4, L1, L1, R1"

	if len(os.Args) < 2 {
		fmt.Println("USEAGE: advent-of-go <day>")
		fmt.Println("<day> = day number of puzzle to run.  e.g. advent-of-go 2")
		return
	}

	partOneResult, partOneErr := day01.PartOne(instructions)
	fmt.Print("Part One ")
	if partOneErr != nil {
		fmt.Print(partOneErr)
	} else {
		fmt.Print(partOneResult)
	}
	partTwoResult, partTwoErr := day01.PartTwo(instructions)

	fmt.Print("Part Two ")
	if partTwoErr != nil {
		fmt.Print(partTwoErr)
	} else {
		fmt.Print(partTwoResult)
	}
}
