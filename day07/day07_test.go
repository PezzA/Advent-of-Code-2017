package day07

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_AbbaDetection(t *testing.T) {
	Convey("Should detect ABBAs", t, func() {
		isAbba, abbaList := isABBA("fgbabbaer")

		So(isAbba, ShouldEqual, true)
		So(abbaList, ShouldContain, "abba")

		isAbba, abbaList = isABBA("wrwqryuuyqerw")
		So(isAbba, ShouldEqual, true)
		So(abbaList, ShouldContain, "yuuy")

		isAbba, abbaList = isABBA("dfgdfgertertdfg")
		So(isAbba, ShouldEqual, false)
		So(abbaList, ShouldHaveLength, 0)

		isAbba, abbaList = isABBA("dfgdfgertertdfguiiu")
		So(isAbba, ShouldEqual, true)
		So(abbaList, ShouldContain, "uiiu")
	})

	Convey("Can detect TLS", t, func() {
		So(supportsTLS("abba[mnop]qrst"), ShouldEqual, true)
		So(supportsTLS("abcd[bddb]xyyx"), ShouldEqual, false)
		So(supportsTLS("aaaa[qwer]tyui"), ShouldEqual, false)
		So(supportsTLS("ioxxoj[asdfgh]zxcvbn"), ShouldEqual, true)
	})

	Convey("Can detect ABAs", t, func() {
		isAba, abaList := isABA("aaaeeetytert")
		So(isAba, ShouldEqual, true)
		So(abaList, ShouldContain, "tyt")

		isAba, abaList = isABA("abaeeetytert")
		So(isAba, ShouldEqual, true)
		So(abaList, ShouldHaveLength, 2)
		So(abaList, ShouldContain, "tyt")
		So(abaList, ShouldContain, "aba")
	})

	Convey("Can reverse aba", t, func() {
		So(reverseABA("aba"), ShouldEqual, "bab")
		So(reverseABA("xyx"), ShouldEqual, "yxy")
	})
}
