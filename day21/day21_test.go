package day21

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Day21_getRotatedPos(t *testing.T) {
	Convey("Get determine where to get a rotated char from", t, func() {
		So(getRotatedPos(3, 10, 1), ShouldEqual, 2)

		So(getRotatedPos(3, 1, 0), ShouldEqual, 1)
		So(getRotatedPos(3, 1, 1), ShouldEqual, 2)
		So(getRotatedPos(3, 1, 2), ShouldEqual, 0)

		So(getRotatedPos(5, 10, 1), ShouldEqual, 1)
		So(getRotatedPos(5, 11, 1), ShouldEqual, 2)
		So(getRotatedPos(5, 12, 1), ShouldEqual, 3)
		So(getRotatedPos(5, 13, 1), ShouldEqual, 4)
		So(getRotatedPos(5, 14, 1), ShouldEqual, 0)
		So(getRotatedPos(5, 15, 1), ShouldEqual, 1)
	})
}

func Test_Day21_rotateLeft(t *testing.T) {
	Convey("Should be able to rotate a string by a specified shift", t, func() {
		var pass password = []rune("wer")
		pass.rotate(1)
		So(pass.toString(), ShouldEqual, "erw")
		pass.rotate(1)
		So(pass.toString(), ShouldEqual, "rwe")
		pass.rotate(1)
		So(pass.toString(), ShouldEqual, "wer")
	})
}

func Test_Day21_rotateRight(t *testing.T) {
	Convey("Should be able to rotate a string by a specified shift", t, func() {
		var pass password = []rune("wer")
		pass.rotate(-1)
		So(pass.toString(), ShouldEqual, "rew")
		pass.rotate(-1)
		So(pass.toString(), ShouldEqual, "wre")
		pass.rotate(-1)
		So(pass.toString(), ShouldEqual, "ewr")
	})
}
func Test_Day21_b(t *testing.T) {
	Convey("Day 21 should be able to", t, func() {
		Convey("Get determine where to get a rotated char from", func() {
			So(getRotatedPos(3, 10, 1), ShouldEqual, 2)

			So(getRotatedPos(3, 1, 0), ShouldEqual, 1)
			So(getRotatedPos(3, 1, 1), ShouldEqual, 2)
			So(getRotatedPos(3, 1, 2), ShouldEqual, 0)

			So(getRotatedPos(5, 10, 1), ShouldEqual, 1)
			So(getRotatedPos(5, 11, 1), ShouldEqual, 2)
			So(getRotatedPos(5, 12, 1), ShouldEqual, 3)
			So(getRotatedPos(5, 13, 1), ShouldEqual, 4)
			So(getRotatedPos(5, 14, 1), ShouldEqual, 0)
			So(getRotatedPos(5, 15, 1), ShouldEqual, 1)
		})

		Convey("Swap character positions in a string", func() {
			var pass password = []rune("wer")
			pass.swapPos(1, 2)
			So(string(pass), ShouldEqual, "wre")
			pass.swapPos(0, 1)
			So(string(pass), ShouldEqual, "rwe")
		})

		Convey("Swap characters in a string", func() {
			var pass password = []rune("wer")
			pass.swapChar(rune('r'), rune('e'))
			So(string(pass), ShouldEqual, "wre")
			pass.swapChar(rune('e'), rune('w'))
			So(string(pass), ShouldEqual, "erw")
		})
	})
}
