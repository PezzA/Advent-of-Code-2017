package day19

// PartOne solves day 20 part one.
func PartOne(input string) (string, error) {

	elves := make([]int, 5)

	counter := 0
	for len(elves) < 1 {

		if counter < len(elves) {

		}

		elves[counter] += elves[counter+1]
	}

	return "", nil
}

func takePresents(elf int, circle []int) []int {
	// take the 1 based index to a zero based index
	elf--

	if elf == len(circle)-1 {
		circle[elf] += circle[0]
		return circle[1:]
	}

	circle[elf] += circle[elf+1]
	return append(circle[:elf+1], circle[elf+2:]...)
}

// PartTwo solves day 20 part one.
func PartTwo(input string) (string, error) {
	return "", nil
}
