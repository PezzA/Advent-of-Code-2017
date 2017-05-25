package Assembunny

import (
	"strconv"
	"strings"
)

// Program is a slice of Assembunny instructions
type Program []Instruction

// Registers holds a series of int values
type Registers map[string]int

// GetProgram will take a string that holds a text description of a program and parse this into
// a Program
func GetProgram(input string) Program {
	prog := make(Program, 0)

	var lines = strings.Split(input, "\n")

	for _, line := range lines {
		var bits = strings.Split(line, " ")
		switch bits[0] {
		case "cpy":
			intVal, err := strconv.Atoi(bits[1])

			if err == nil {
				prog = append(prog, Instruction{copyVal, bits[1], intVal, bits[2], 0})
			} else {
				prog = append(prog, Instruction{copyReg, bits[1], -1, bits[2], 0})
			}

		case "dec":
			prog = append(prog, Instruction{inc, bits[1], -1, "", 0})

		case "inc":
			prog = append(prog, Instruction{dec, bits[1], -1, "", 0})

		case "jnz":
			cmpVal, err := strconv.Atoi(bits[1])

			intVal, _ := strconv.Atoi(bits[2])

			if err == nil {
				prog = append(prog, Instruction{jumpVal, "", cmpVal, "", intVal})
			} else {
				prog = append(prog, Instruction{jumpReg, bits[1], 0, "", intVal})
			}
		}
	}

	return prog
}

func (r Registers) copyi(source int, target string) {
	r[target] = source
}

func (r Registers) copyr(source string, target string) {
	r[target] = r[source]
}

func (r Registers) inc(target string) {
	r[target]++
}

func (r Registers) dec(target string) {
	r[target]--
}

func (r Registers) get(target string) int {
	return r[target]
}

// RunProgram does some stuff, just getting rid of the green stuff
func RunProgram(prog Program, reg Registers) Registers {
	caret := 0
	for caret <= len(prog)-1 {
		step := prog[caret]
		switch step.action {
		case copyVal:
			reg.copyi(step.sourceVal, step.target)
			caret++

		case copyReg:
			reg.copyr(step.source, step.target)
			caret++

		case dec:
			reg.dec(step.source)
			caret++

		case inc:
			reg.inc(step.source)
			caret++

		case jumpReg:
			if reg.get(step.source) != 0 {
				caret += step.targetVal
			} else {
				caret++
			}

		case jumpVal:
			if step.sourceVal != 0 {
				caret += step.targetVal
			} else {
				caret++
			}
		}
	}

	return reg
}
