package day12

import (
	"testing"

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
			prog := getProgram(testInput())
			So(len(prog), ShouldEqual, 6)
		})

		Convey("Run the test program", func() {
			prog := getProgram(testInput())
			registers := makeRegisters()

			result := runProgram(prog, registers, "a")

			So(result, ShouldEqual, 42)
		})

		Convey("Run the real program", func() {
			prog := getProgram(PuzzleInput())
			registers := makeRegisters()

			result := runProgram(prog, registers, "a")
			So(result, ShouldEqual, 318020)
		})
	})

}
