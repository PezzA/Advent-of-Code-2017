package day18

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func testInput() string {
	return `.^^.^.^^^^`
}

func Test_Day18(t *testing.T) {
	Convey("Day 18 should be able to", t, func() {
		Convey("-> Detect traps!", func() {
			So(isTrap(trap, trap, safeSpace), ShouldEqual, trap)
			So(isTrap(safeSpace, trap, trap), ShouldEqual, trap)
			So(isTrap(empty, empty, trap), ShouldEqual, trap)
			So(isTrap(empty, safeSpace, trap), ShouldEqual, trap)
			So(isTrap(trap, safeSpace, safeSpace), ShouldEqual, trap)
		})
		Convey("-> And detect safe spaces!", func() {
			So(isTrap(safeSpace, safeSpace, safeSpace), ShouldEqual, safeSpace)
			So(isTrap(empty, empty, empty), ShouldEqual, safeSpace)
			So(isTrap(trap, trap, trap), ShouldEqual, safeSpace)
			So(isTrap(trap, safeSpace, trap), ShouldEqual, safeSpace)
		})

		Convey("-> Should be able to process a row", func() {
			So(processRow([]rune("..^^.")), ShouldResemble, []rune(".^^^^"))
		})

		Convey("-> Should be able to count space types in a row", func() {
			So(countType([]rune("..^^."), safeSpace), ShouldEqual, 3)
			So(countType([]rune("..^^."), trap), ShouldEqual, 2)
			So(countType([]rune("....."), trap), ShouldEqual, 0)
			So(countType([]rune("....."), safeSpace), ShouldEqual, 5)
			So(countType([]rune(""), safeSpace), ShouldEqual, 0)
		})

		/*	Commented due to the puzzle having to run for a specific amount of iterations, which sort of breaks the string puzzle input paradigm.  could add it as a csv?
			Convey("-> Should get the test data right", func() {
			result, error := PartOne(testInput())
			So(result, ShouldEqual, "38")
			So(error, ShouldEqual, nil)
		})*/
	})
}
