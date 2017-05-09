package day19

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Day19(t *testing.T) {

	Convey("Should be able to determine the adjacent elf", t, func() {
		So(getAdjacentElf(4, 1), ShouldEqual, 3)
		So(getAdjacentElf(4, 2), ShouldEqual, 4)
		So(getAdjacentElf(4, 4), ShouldEqual, 2)
		So(getAdjacentElf(5, 1), ShouldEqual, 3)
		So(getAdjacentElf(5, 2), ShouldEqual, 4)
	})

	Convey("Should be able to determine the next active elf", t, func() {
		elves := circle{0, 0, 1, 1, 1}

		So(elves.getNextActiveElf(0), ShouldEqual, 2)
		So(elves.getNextActiveElf(2), ShouldEqual, 3)
		So(elves.getNextActiveElf(4), ShouldEqual, 2)
		So(elves.getNextActiveElf(3), ShouldEqual, 4)
	})

	Convey("Should be able to resolve a circle", t, func() {
		elves := circle{1, 1, 1, 1, 1}

		nextElf := 0
		finished := false

		finished = elves.takePresentsFromLeft(nextElf)
		So(elves, ShouldResemble, circle{2, 0, 1, 1, 1})
		So(finished, ShouldBeFalse)

		nextElf = elves.getNextActiveElf(nextElf)
		So(nextElf, ShouldEqual, 2)

		finished = elves.takePresentsFromLeft(nextElf)
		So(elves, ShouldResemble, circle{2, 0, 2, 0, 1})
		So(finished, ShouldBeFalse)

		nextElf = elves.getNextActiveElf(nextElf)
		So(nextElf, ShouldEqual, 4)

		finished = elves.takePresentsFromLeft(nextElf)
		So(elves, ShouldResemble, circle{0, 0, 2, 0, 3})
		So(finished, ShouldBeFalse)

		nextElf = elves.getNextActiveElf(nextElf)
		So(nextElf, ShouldEqual, 2)

		finished = elves.takePresentsFromLeft(nextElf)
		So(elves, ShouldResemble, circle{0, 0, 5, 0, 0})
		So(finished, ShouldBeTrue)
		So(nextElf, ShouldEqual, 2)

	})

}
