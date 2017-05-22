package day12

import (
	"testing"

	"log"

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
			log.Println(getProgram(testInput()))
		})
	})
}
