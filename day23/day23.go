package day23

import (
	"strconv"

	"github.com/pezza/advent-of-go/Assembunny"
)

func makeRegisters() Assembunny.Registers {
	reg := make(Assembunny.Registers)

	reg["a"] = 0
	reg["b"] = 0
	reg["c"] = 0
	reg["d"] = 0

	return reg
}

// PartOne solves day 24 part one.
func PartOne(input string) (string, error) {
	prog := Assembunny.GetProgram(PuzzleInput())
	registers := makeRegisters()

	registers = Assembunny.RunProgram(prog, registers)

	return strconv.Itoa(registers["a"]), nil
}

// PartTwo solves day 24 part one.
func PartTwo(input string) (string, error) {
	return "", nil
}
