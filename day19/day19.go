package day19

import (
	"log"
)

type elf struct {
	index    int
	presents int
}

func getElves(count int) []elf {
	elves := make([]elf, count)

	for index := range elves {
		elves[index] = elf{index + 1, 1}
	}

	return elves
}

// PartOne solves day 20 part one.
func PartOne(input string) (string, error) {
	elves := getElves(5)
	position := 1

	for len(elves) > 1 {
		elves = takePresents(position, elves)

		if position < len(elves) {
			position++
		} else {
			position = 1
		}

		log.Println(position)
	}

	return "", nil
}

func takePresents(elf int, circle []elf) []elf {
	indexElf := elf - 1

	if elf == len(circle) {
		circle[indexElf].presents += circle[0].presents
		return circle[1:]
	}

	circle[indexElf].presents += circle[indexElf+1].presents

	return circle
}

// PartTwo solves day 20 part one.
func PartTwo(input string) (string, error) {
	return "", nil
}
