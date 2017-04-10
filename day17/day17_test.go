package day17

import "testing"
import . "github.com/smartystreets/goconvey/convey"
import "log"

func Test_Day17(t *testing.T) {
	Convey("-> Should be able to do some test stuff", t, func() {
		log.Println(getHash(testInput()))

		log.Println(getDoors(PuzzleInput(), point{0, 0}))
	})

	Convey("-> Should be able to tell if a door is open or not", t, func() {
		So(isOpen('a'), ShouldBeFalse)
		So(isOpen('b'), ShouldBeTrue)
		So(isOpen('c'), ShouldBeTrue)
		So(isOpen('d'), ShouldBeTrue)
		So(isOpen('e'), ShouldBeTrue)
		So(isOpen('f'), ShouldBeTrue)
		So(isOpen('g'), ShouldBeFalse)
		So(isOpen('h'), ShouldBeFalse)
		So(isOpen('i'), ShouldBeFalse)
		So(isOpen('j'), ShouldBeFalse)
		So(isOpen('k'), ShouldBeFalse)
	})
}
