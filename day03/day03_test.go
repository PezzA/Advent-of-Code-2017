package day03

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func getTestData() string {
	return `101 301 501
  102  302 502
  303 103		503
  201  401  601
  20  602  102
  603 403  103`
}

func Test_Loading_And_Parsing_Columns(t *testing.T) {
	triangles := getColTriangles(getTestData())
	Convey("After parsing the column triangle list", t, func() {
		Convey("It should have 6 triangles", func() {
			So(triangles, ShouldHaveLength, 6)
		})

		Convey("It should be fine with tabs", func() {
			So(triangles[2].side1, ShouldEqual, 501)
			So(triangles[2].side2, ShouldEqual, 502)
			So(triangles[2].side3, ShouldEqual, 503)
		})

		Convey("It should be fine with excess spaces", func() {
			So(triangles[5].side1, ShouldEqual, 601)
			So(triangles[5].side2, ShouldEqual, 102)
			So(triangles[5].side3, ShouldEqual, 103)
		})
	})
}

func Test_Loading_And_Parsing_The_File(t *testing.T) {
	triangles := getTriangles(getTestData())

	Convey("After parsing the test triangle list", t, func() {
		Convey("It should have 6 triangles", func() {
			So(triangles, ShouldHaveLength, 6)
		})

		Convey("It should be fine with tabs", func() {
			So(triangles[2].side1, ShouldEqual, 303)
			So(triangles[2].side2, ShouldEqual, 103)
			So(triangles[2].side3, ShouldEqual, 503)
		})

		Convey("It should be fine with excess spaces", func() {
			So(triangles[5].side1, ShouldEqual, 603)
			So(triangles[5].side2, ShouldEqual, 403)
			So(triangles[5].side3, ShouldEqual, 103)
		})
	})
}

func Test_Testing_the_Triangle(t *testing.T) {
	t1 := triangle{5, 10, 25}
	t2 := triangle{5, 10, 15}

	Convey("When examining a triangle ", t, func() {
		Convey("It should detect t1 as invalid", func() {
			So(t1.isValid(), ShouldBeFalse)
		})
		Convey("It should detect t2 as invalid", func() {
			So(t2.isValid(), ShouldBeFalse)
		})
	})
}
