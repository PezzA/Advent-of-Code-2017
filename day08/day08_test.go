package day08

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Regex(t *testing.T) {
	Convey("Instruction Regex", t, func() {
		Convey("Should detect rect instructions", func() {
			instruction, error := parseInstruction("rect 2x2")

			So(error, ShouldEqual, nil)
			So(instruction.instruction, ShouldEqual, "")
			So(instruction.reg1, ShouldEqual, 2)
			So(instruction.reg2, ShouldEqual, 2)
		})
	})
}
