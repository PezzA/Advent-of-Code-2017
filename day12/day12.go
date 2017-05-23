package day12

import (
	"strconv"
	"strings"
)

type registers map[string]int

type instruction struct {
	action    int
	source    string
	sourceVal int
	target    string
}

type program []instruction

func makeRegisters() registers {
	reg := make(registers)

	reg["a"] = 0
	reg["b"] = 0
	reg["c"] = 0
	reg["d"] = 0

	return reg
}

func getProgram(input string) program {
	prog := make(program, 0)

	var lines = strings.Split(input, "\n")

	for _, line := range lines {
		var bits = strings.Split(line, " ")
		switch bits[0] {
		case "cpy":
			intVal, err := strconv.Atoi(bits[1])

			if err == nil {
				prog = append(prog, instruction{1, bits[1], intVal, bits[2]})
			} else {
				prog = append(prog, instruction{2, bits[1], -1, bits[2]})
			}

		case "dec":
			prog = append(prog, instruction{3, bits[1], -1, ""})
		case "inc":
			prog = append(prog, instruction{4, bits[1], -1, ""})
		case "jnz":
			intVal, _ := strconv.Atoi(bits[2])
			prog = append(prog, instruction{5, bits[1], intVal, ""})
		}
	}

	return prog
}

func (r registers) copyi(source int, target string) {
	r[target] = source
}

func (r registers) copyr(source string, target string) {
	r[target] = r[source]
}

func (r registers) inc(target string) {
	r[target]++
}

func (r registers) dec(target string) {
	r[target]--
}

func (r registers) get(target string) int {
	return r[target]
}

func runProgram(prog program, reg registers, finalPin string) int {
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
				caret += step.sourceVal
			} else {
				caret++
			}
		}

	}

	return reg.get("a")
}

// PartOne solves day 12 part one.
func PartOne(input string) (string, error) {
	return "", nil
}

// PartTwo solves day 12 part one.
func PartTwo(input string) (string, error) {
	return "", nil
}
