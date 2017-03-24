package day16

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_day16(t *testing.T) {
	Convey("It should do something", t, func() {
		Convey("Such as successfully expand a string", func() {
			So(expand(stringToBool("1")), ShouldResemble, stringToBool("100"))
			So(expand(stringToBool("0")), ShouldResemble, stringToBool("001"))
			So(expand(stringToBool("11111")), ShouldResemble, stringToBool("11111000000"))
			So(expand(stringToBool("111100001010")), ShouldResemble, stringToBool("1111000010100101011110000"))
		})

		Convey("Should successfully reduce a string", func() {
			So(reduce(stringToBool("110010110100")), ShouldResemble, stringToBool("110101"))
		})

		Convey("Should be able to calculate a checksum", func() {
			So(checkSum(stringToBool("110010110100")), ShouldResemble, stringToBool("100"))
		})
	})
}
