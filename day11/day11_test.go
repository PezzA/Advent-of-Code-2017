package day11

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Day11(t *testing.T) {
	Convey("Day 11 should be able to ", t, func() {
		f := getStartingFacility(testInput())

		Convey("-> Parse the puzzle input", func() {
			components := make([]component, 0)
			components = append(components, component{0, "hydrogen", true})
			components = append(components, component{0, "lithium", true})
			components = append(components, component{1, "hydrogen", false})
			components = append(components, component{2, "lithium", false})

			So(f, ShouldResemble, facility{0, 4, components})
		})

		/*
			Convey("-> Draw the facility", func() {
				f.drawFacility()
			})
		*/

		Convey("-> Get all the stuff on the specified floor", func() {
			So(f.getItemsOnFloor(0)[0], ShouldResemble, component{0, "hydrogen", true})
			So(f.getItemsOnFloor(0)[1], ShouldResemble, component{0, "lithium", true})
			So(f.getItemsOnFloor(1)[0], ShouldResemble, component{1, "hydrogen", false})
			So(f.getItemsOnFloor(2)[0], ShouldResemble, component{2, "lithium", false})
			So(len(f.getItemsOnFloor(3)), ShouldResemble, 0)
		})

		/*
			Convey("-> Give all the cominbinations for the stuff on the floor", func() {
				items := elevatorCombinations(f.getItemsOnFloor(0))
				So(items[0], ShouldResemble, []string{"HM"})
				So(items[1], ShouldResemble, []string{"LM"})
				So(items[2], ShouldResemble, []string{"HM", "LM"})
			})

			Convey("-> Should be able to detect if moves are valid", func() {

				So(f.validMove([]string{"LM"}, 1), ShouldEqual, false)
				So(f.validMove([]string{"HM"}, 1), ShouldEqual, true)
				So(f.validMove([]string{"HM", "LM"}, 1), ShouldEqual, false)

				f.elevator = 1

				So(f.validMove([]string{"HG"}, 0), ShouldEqual, false)
				So(f.validMove([]string{"HG"}, 2), ShouldEqual, true)

			})

		*/
	})
}
