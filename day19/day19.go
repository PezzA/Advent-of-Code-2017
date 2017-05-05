package day19

import (
	"log"
)

type elf struct {
	index    int
	presents int
}

type circle []elf

func (c circle) getNextActiveElf(currentPosition int) int {
	if currentPosition == len(c) {
		currentPosition = 0
	}

	currentElf := c[currentPosition]

	for currentElf.presents == 0 {
		if currentPosition < len(c)-1 {
			currentPosition++
		} else {
			currentPosition = 0
		}
		currentElf = c[currentPosition]
	}

	return currentElf.index
}

func getElves(count int) circle {
	elves := make([]elf, count)

	for index := range elves {
		elves[index] = elf{index + 1, 1}
	}

	return elves
}

func (c circle) takePresents(circlePosition int) {
	nextElf := c.getNextActiveElf(circlePosition - 1)
	log.Println(nextElf)
	c[circlePosition-1].presents += c[nextElf].presents
	c[nextElf].presents = 0
}

// PartOne solves day 20 part one.
func PartOne(input string) (string, error) {
	elves := getElves(5)

	circlePosition := 1

	loops := 0
	for {
		if loops == 10 {
			break
		}
		elves.takePresents(circlePosition)

		log.Println(circlePosition, elves)

		if elves[circlePosition-1].presents == 5 {
			break
		}

		circlePosition = elves.getNextActiveElf(circlePosition)

		loops++
	}

	return "", nil
}

// PartTwo solves day 20 part one.
func PartTwo(input string) (string, error) {
	return "", nil
}
