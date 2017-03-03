package day04

import (
	"errors"
	"sort"
	"strconv"
	"strings"

	. "github.com/pezza/advent-of-go/mapSorter"
)

type roomData string

func PartOne() (string, error) {
	overall := 0
	for _, line := range strings.Split(PuzzleInput(), "\n") {
		room := roomData(line)

		if room.isValid() {
			overall += room.getSectorID()
		}
	}

	return strconv.Itoa(overall), nil
}

func PartTwo() (string, error) {
	for _, line := range strings.Split(PuzzleInput(), "\n") {

		room := roomData(line)

		if room.decryptRoomName() == "northpole object storage " {
			return strconv.Itoa(room.getSectorID()), nil
		}

	}

	return "", errors.New("Could not find room")
}

func rotateRune(r rune, shift int) rune {
	actualShift := shift % 26

	newRune := int(r) + actualShift
	if newRune < 123 {
		return rune(newRune)
	}

	return rune(newRune - 26)
}

func (r roomData) getBits() []string {
	var data = string(r)

	return strings.Split(data, "-")
}

func (r roomData) decryptRoomName() string {
	bits := r.getBits()
	sectorID := r.getSectorID()

	var name string
	for x, line := range bits {
		if x != len(bits)-1 {
			var word string
			for _, bit := range line {
				word += string(rotateRune(bit, sectorID))
			}

			name += word + " "
		}
	}

	return name
}

func (r roomData) getSectorID() int {
	var bits = r.getBits()
	var lastBit = bits[len(bits)-1]

	textualSectorID := strings.Split(lastBit, "[")[0]

	sectorID, _ := strconv.Atoi(textualSectorID)

	return sectorID
}

func (r roomData) getCheckSum() string {
	var bits = r.getBits()

	alphaMap := make(map[rune]int)

	for x, line := range bits {
		if x != len(bits)-1 {
			for _, bit := range line {
				if _, ok := alphaMap[bit]; ok {
					alphaMap[bit]++
				} else {
					alphaMap[bit] = 1
				}
			}
		}
	}

	sortedMap := MapToList(alphaMap)

	sort.Sort(sort.Reverse(sortedMap))

	checkSum := ""

	for _, val := range sortedMap[:5] {
		checkSum += string(val.Key)
	}

	return checkSum
}

func (r roomData) isValid() bool {
	return r.getOrigCheckSum() == r.getCheckSum()
}

func (r roomData) getOrigCheckSum() string {
	var bits = r.getBits()
	var lastBit = bits[len(bits)-1]

	checkSum := strings.Split(strings.Split(lastBit, "[")[1], "]")[0]
	return checkSum
}

func getAlphaMap(s string) map[rune]int {
	alphaMap := make(map[rune]int)

	return alphaMap
}
