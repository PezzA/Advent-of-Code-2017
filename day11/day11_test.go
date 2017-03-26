package day11

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Day11(t *testing.T) {
	Convey("Day 11 should be able to ", t, func() {

		f := getStartingFacility(testInput())

		Convey("-> Parse the puzzle input", func() {
			reactorList := make(reactorList)
			reactorList["hydrogen"] = reactor{0, 1}
			reactorList["lithium"] = reactor{0, 2}

			elementList := []string{"hydrogen", "lithium"}
			So(f, ShouldResemble, facility{0, 4, elementList, reactorList})
		})

		Convey("-> Draw the facility", func() {
			f.drawFacility()
		})
	})
}
