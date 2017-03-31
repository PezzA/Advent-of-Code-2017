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

		Convey("-> Get all the stuff on the specified floor", func() {
			So(f.getItemsOnFloor(0), ShouldResemble, []string{"HM", "LM"})
			So(f.getItemsOnFloor(1), ShouldResemble, []string{"HG"})
			So(f.getItemsOnFloor(2), ShouldResemble, []string{"LG"})
			So(len(f.getItemsOnFloor(3)), ShouldResemble, 0)
		})

		Convey("-> Give all the cominbinations for the stuff on the floor", func() {
			items := elevatorCombinations(f.getItemsOnFloor(0))
			So(items[0], ShouldResemble, []string{"HM"})
			So(items[1], ShouldResemble, []string{"LM"})
			So(items[2], ShouldResemble, []string{"HM", "LM"})
		})
	})
}
