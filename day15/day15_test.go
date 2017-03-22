package day15

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_day15(t *testing.T) {

	Convey("It should do something", t, func() {
		Convey("like being able to tell what position a disc will be at, at a given time", func() {
			disc1 := disc{5, 4}
			So(disc1.GetTimeLaspedPosition(0), ShouldEqual, 4)
			So(disc1.GetTimeLaspedPosition(1), ShouldEqual, 0)
			So(disc1.GetTimeLaspedPosition(2), ShouldEqual, 1)
			So(disc1.GetTimeLaspedPosition(3), ShouldEqual, 2)
			So(disc1.GetTimeLaspedPosition(4), ShouldEqual, 3)
			So(disc1.GetTimeLaspedPosition(5), ShouldEqual, 4)
			So(disc1.GetTimeLaspedPosition(6), ShouldEqual, 0)
		})
	})
}
