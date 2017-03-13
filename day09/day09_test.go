package day09

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Regex(t *testing.T) {

	Convey("Should be able to split a string into its sections", t, func() {
		sections := getAllSections("A(1x5)BC")

		So(sections, ShouldHaveLength, 3)
		So(sections[0], ShouldResemble, section{"A", 1})
		So(sections[1], ShouldResemble, section{"B", 5})
		So(sections[2], ShouldResemble, section{"C", 1})

		Convey("and be able to predict its length", func() {
			So(getRenderedLength(sections), ShouldEqual, 7)
		})
	})

	Convey("Should be able to fully decompress a section", t, func() {
		So(getLength("(27x12)(20x12)(13x14)(7x10)(1x12)A"), ShouldEqual, 241920)
		So(getLength("X(8x2)(3x3)ABCY"), ShouldEqual, 20)
		So(getLength("(25x3)(3x3)ABC(2x3)XY(5x2)PQRSTX(18x9)(3x2)TWO(5x7)SEVEN"), ShouldEqual, 445)
	})
}
