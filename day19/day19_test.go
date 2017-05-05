package day19

import (
	"testing"

	"log"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Day19(t *testing.T) {

	Convey("Should be able to create elves", t, func() {
		log.Println(getElves(5))
	})

	Convey("Should be able to take all the presents of an elf", t, func() {
		elves := []elf{elf{1, 1}, elf{2, 1}, elf{3, 1}, elf{4, 1}, elf{5, 1}}

		elves = takePresents(1, elves)
		So(elves, ShouldResemble, []elf{elf{1, 2}, elf{3, 1}, elf{4, 1}, elf{5, 1}})

		elves = takePresents(4, elves)
		So(elves, ShouldResemble, []elf{elf{3, 1}, elf{4, 1}, elf{5, 3}})
	})
}
