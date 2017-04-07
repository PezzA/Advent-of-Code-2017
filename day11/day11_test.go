package day11

import (
	"testing"

	"log"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Day11(t *testing.T) {
	Convey("Day 11 should be able to ", t, func() {
		f := getStartingFacility(testInput())

		hydrogenMicrochip := component{0, "hydrogen", true}
		lithiumMicrochip := component{0, "lithium", true}
		hydrogenGenerator := component{1, "hydrogen", false}
		lithiumGenerator := component{2, "lithium", false}

		Convey("-> Parse the puzzle input", func() {
			components := make([]component, 0)
			components = append(components, hydrogenMicrochip)
			components = append(components, lithiumMicrochip)
			components = append(components, hydrogenGenerator)
			components = append(components, lithiumGenerator)

			So(f, ShouldResemble, facility{0, components})
		})

		Convey("-> Draw the facility", func() {
			f.drawFacility()
		})

		Convey("-> Get all the stuff on the specified floor", func() {
			So(f.getItemsOnFloor(0)[0], ShouldResemble, hydrogenMicrochip)
			So(f.getItemsOnFloor(0)[1], ShouldResemble, lithiumMicrochip)
			So(f.getItemsOnFloor(1)[0], ShouldResemble, hydrogenGenerator)
			So(f.getItemsOnFloor(2)[0], ShouldResemble, lithiumGenerator)
			So(len(f.getItemsOnFloor(3)), ShouldResemble, 0)
		})

		Convey("-> Give all the cominbinations for the stuff on the floor", func() {
			items := elevatorCombinations(f.getItemsOnFloor(0))
			So(items[0], ShouldResemble, []component{hydrogenMicrochip})
			So(items[1], ShouldResemble, []component{lithiumMicrochip})
			So(items[2], ShouldResemble, []component{hydrogenMicrochip, lithiumMicrochip})
		})

		Convey("-> Should be able to detect if moves are valid", func() {

			log.Println("starting facility", f)
			state := solve(f)
			for _, newstate := range state.subStates {
				log.Println("Solving Substate", newstate.f)
				log.Println(solve(newstate.f))
			}
		})

	})
}
