package day08

import "regexp"
import "log"

var instructionRegex = regexp.MustCompile("(rect |rotate column x=|rotate row y=)([0-9]*)(x| by )([0-9]*)")

type instruction struct {
	instruction string
	reg1        int
	reg2        int
}

func parseInstruction(input string) (instruction, error) {
	match := instructionRegex.FindStringSubmatch(input)

	log.Println(match)

	return instruction{}, nil
}

// PartOne executes day08 part one!
func PartOne() (string, error) {
	return "", nil
}

// PartTwo executes day08 part two!
func PartTwo() (string, error) {
	return "", nil
}
