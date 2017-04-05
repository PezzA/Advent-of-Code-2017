package day11

import (
	"math"
	"strconv"
	"strings"
)

// PartOne returns the solution to day11 part one
func PartOne(input string) (string, error) {
	return "", nil
}

// PartTwo returns the solution to day11 part two
func PartTwo(input string) (string, error) {
	return "", nil
}

func (f facility) validMove(bits []string, floor int) bool {
	allBits := append(bits, f.getItemsOnFloor(floor)...)
	counts := make(map[string]int)

	// if we have 2 of the same bits, all good, if we have just one, bad
	for _, bit := range allBits {
		initial := bit[:1]
		if _, ok := counts[initial]; ok {
			counts[initial]++
		} else {
			counts[initial] = 1
		}
	}

	for _, v := range counts {
		if v == 1 {
			return false
		}
	}

	return true
}

// get all the items on the specified floor
func (f facility) getItemsOnFloor(floor int) []string {
	bits := make([]string, 0)
	for k, v := range f.reactors {
		if v.generator == floor {
			bits = append(bits, strings.ToUpper(k[:1]+"G"))
		}
		if v.microchip == floor {
			bits = append(bits, strings.ToUpper(k[:1]+"M"))
		}
	}
	return bits
}

// Given a list of items, returns a list of
// all the possibly combinations of items
// that could be put into the elevator.
func elevatorCombinations(bits []string) [][]string {
	parts := make([][]string, 0)
	maxCard := int(math.Pow(2, float64(len(bits))))

	for i := 0; i < maxCard; i++ {
		takeParts := make([]string, 0)
		count := strings.Count(strconv.FormatInt(int64(i), 2), "1")
		if count == 1 || count == 2 {
			for j := 0; j < len(bits); j++ {
				if i&(1<<uint(j)) != 0 {
					takeParts = append(takeParts, bits[j])
				}
			}
		}
		if len(takeParts) > 0 {
			parts = append(parts, takeParts)
		}
	}
	return parts
}
