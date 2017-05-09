package day19

import (
	"strconv"
)

type circle []int

func getAdjacentElf(circleSize int, position int) int {
	return (circleSize / 2) + position ^ circleSize
}

// getNextActiveElf uses zero based index
func (c circle) getNextActiveElf(currentPosition int) int {
	if currentPosition == len(c)-1 {
		currentPosition = 0
	} else {
		currentPosition++
	}

	currentElf := c[currentPosition]

	for currentElf == 0 {
		if currentPosition < len(c)-1 {
			currentPosition++
		} else {
			currentPosition = 0
		}
		currentElf = c[currentPosition]
	}

	return currentPosition
}

func getElves(count int) circle {
	elves := make([]int, count)

	for index := range elves {
		elves[index] = 1
	}

	return elves
}

func (c circle) takePresentsFromLeft(circlePosition int) bool {
	nextElf := c.getNextActiveElf(circlePosition)
	c[circlePosition] += c[nextElf]
	c[nextElf] = 0

	return c[circlePosition] == len(c)
}

// PartOne solves day 20 part one.
func PartOne(input string) (string, error) {
	circleSize, _ := strconv.Atoi(input)

	elves := getElves(circleSize)

	nextElf := 0
	finished := false

	for !finished {
		finished = elves.takePresentsFromLeft(nextElf)
		nextElf = elves.getNextActiveElf(nextElf)
	}
	return strconv.Itoa(nextElf + 1), nil
}

// PartTwo solves day 20 part one.
func PartTwo(input string) (string, error) {
	circleSize, _ := strconv.Atoi(input)

	elves := getElves(circleSize)

	nextElf := 0
	finished := false

	for !finished {
		finished = elves.takePresentsFromLeft(nextElf)
		nextElf = elves.getNextActiveElf(nextElf)
	}
	return strconv.Itoa(nextElf + 1), nil
}
