package day12

import (
	"testing"

	"github.com/pezza/advent-of-go/Assembunny"
	. "github.com/smartystreets/goconvey/convey"
)

func testInput() string {
	return `cpy 41 a
inc a
inc a
dec a
jnz a 2
dec a`
}

func Test_Day12(t *testing.T) {
	Convey("Day 12 should be able to", t, func() {
		Convey("Load a program", func() {
			prog := Assembunny.GetProgram(testInput())
			So(len(prog), ShouldEqual, 6)
		})

		Convey("Run the test program", func() {
			prog := Assembunny.GetProgram(testInput())
			registers := makeRegisters()

			registers = Assembunny.RunProgram(prog, registers)

			So(registers["a"], ShouldEqual, 42)
		})

		Convey("Run the real program", func() {
			prog := Assembunny.GetProgram(PuzzleInput())
			registers := makeRegisters()

			registers = Assembunny.RunProgram(prog, registers)
			So(registers["a"], ShouldEqual, 318020)
		})

		Convey("Run the modified real program", func() {
			prog := Assembunny.GetProgram(PuzzleInput())
			registers := makeRegisters()

			registers["c"] = 1

			registers = Assembunny.RunProgram(prog, registers)
			So(registers["a"], ShouldEqual, 9227674)
		})
	})

}
