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

	Convey("Should be able to determine the next active elf", t, func() {
		elves := circle{elf{1, 0}, elf{2, 0}, elf{3, 1}, elf{4, 1}, elf{5, 1}}

		So(elves.getNextActiveElf(1), ShouldEqual, 3)
		So(elves.getNextActiveElf(3), ShouldEqual, 4)
		So(elves.getNextActiveElf(5), ShouldEqual, 3)
		So(elves.getNextActiveElf(4), ShouldEqual, 5)
	})

	Convey("Should be able to take all the presents of an elf", t, func() {
		elves := circle{elf{1, 1}, elf{2, 1}, elf{3, 1}, elf{4, 1}, elf{5, 1}}

		elves.takePresents(1)
		So(elves, ShouldResemble, circle{elf{1, 2}, elf{2, 0}, elf{3, 1}, elf{4, 1}, elf{5, 1}})

		elves = circle{elf{1, 0}, elf{2, 0}, elf{3, 2}, elf{4, 0}, elf{5, 3}}
		elves.takePresents(3)
		So(elves, ShouldResemble, circle{elf{1, 0}, elf{2, 0}, elf{3, 5}, elf{4, 0}, elf{5, 0}})
	})

}
