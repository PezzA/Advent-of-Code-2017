package day19

import (
	"testing"

	"log"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Day19(t *testing.T) {
	Convey("Should be able to take all the presents of an elf", t, func() {
		elves := []int{1, 1, 1, 1, 1}
		log.Println()

		elves = takePresents(1, elves)
		So(elves, ShouldResemble, []int{2, 1, 1, 1})
		elves = takePresents(4, elves)
		So(elves, ShouldResemble, []int{1, 1, 3})
	})
}
