package Assembunny

import (
	"strconv"
	"strings"
)

type Instruction struct {
	action    int
	source    string
	sourceVal int
	target    string
	targetVal int
}

type Program []Instruction

type Registers map[string]int

func GetProgram(input string) Program {
	prog := make(Program, 0)

	var lines = strings.Split(input, "\n")

	for _, line := range lines {
		var bits = strings.Split(line, " ")
		switch bits[0] {
		case "cpy":
			intVal, err := strconv.Atoi(bits[1])

			if err == nil {
				prog = append(prog, Instruction{1, bits[1], intVal, bits[2], 0})
			} else {
				prog = append(prog, Instruction{2, bits[1], -1, bits[2], 0})
			}

		case "dec":
			prog = append(prog, Instruction{3, bits[1], -1, "", 0})
		case "inc":
			prog = append(prog, Instruction{4, bits[1], -1, "", 0})
		case "jnz":
			cmpVal, err := strconv.Atoi(bits[1])

			intVal, _ := strconv.Atoi(bits[2])

			if err == nil {
				prog = append(prog, Instruction{6, "", cmpVal, "", intVal})
			} else {
				prog = append(prog, Instruction{5, bits[1], 0, "", intVal})
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

func RunProgram(prog Program, reg Registers, finalPin string) int {
	caret := 0

	for caret <= len(prog)-1 {
		step := prog[caret]
		switch step.action {
		case 1:
			reg.copyi(step.sourceVal, step.target)
			caret++
		case 2:
			reg.copyr(step.source, step.target)
			caret++
		case 3:
			reg.dec(step.source)
			caret++
		case 4:
			reg.inc(step.source)
			caret++
		case 5:
			if reg.get(step.source) != 0 {
				caret += step.targetVal
			} else {
				caret++
			}
		case 6:
			if step.sourceVal != 0 {
				caret += step.targetVal
			} else {
				caret++
			}
		}

	}

	return reg.get("a")
}
