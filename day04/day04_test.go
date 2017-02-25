package day04

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRuneRotator(t *testing.T) {
	Convey("Rotating Runes", t, func() {
		Convey("should work!!", func() {
			So(rotateRune('a', 26), ShouldEqual, 'a')
			So(rotateRune('a', 27), ShouldEqual, 'b')
			So(rotateRune('a', 1), ShouldEqual, 'b')
		})
	})
}

func TestRoomListParsing(t *testing.T) {
	var room roomData
	room = "a-b-c-d-e-f-g-h-987[abcde]"

	Convey("After setting up a roomlist", t, func() {

		Convey("It should be able to return the SectorId", func() {
			So(room.getSectorID(), ShouldEqual, 987)
		})

		Convey("It should be able to parse out the given checksum", func() {
			So(room.getOrigCheckSum(), ShouldEqual, "abcde")
		})

		Convey("It should be able to confirm the checksum is correct", func() {
			So(room.isValid(), ShouldEqual, true)
		})

		Convey("It should be able to confirm the checksum is not valid", func() {
			var badRoom roomData = "totally-real-room-200[decoy]"
			So(badRoom.isValid(), ShouldEqual, false)
		})
	})
}
