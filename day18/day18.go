package day18

import "strconv"

const (
	trap      = 94
	safeSpace = 46
	empty     = 0
)

func isTrap(left rune, center rune, right rune) rune {
	if (left == trap && center == trap && right != trap) ||
		(left != trap && center == trap && right == trap) ||
		(left != trap && center != trap && right == trap) ||
		(left == trap && center != trap && right != trap) {
		return trap
	}
	return safeSpace
}

func processRow(input []rune) []rune {
	row := make([]rune, 0)

	for i := 0; i < len(input); i++ {
		switch i {
		case 0:
			row = append(row, isTrap(empty, input[i], input[i+1]))
		case len(input) - 1:
			row = append(row, isTrap(input[i-1], input[i], empty))
		default:
			row = append(row, isTrap(input[i-1], input[i], input[i+1]))
		}
	}
	return row
}

func countType(row []rune, spaceType rune) int {
	count := 0
	for _, space := range row {
		if space == spaceType {
			count++
		}
	}
	return count
}

// PartOne solves day 18 part one.
func PartOne(input string) (string, error) {
	currentRow := []rune(input)
	safeSpaceCount := countType(currentRow, safeSpace)

	for i := 0; i < 39; i++ {
		currentRow = processRow(currentRow)
		safeSpaceCount += countType(currentRow, safeSpace)
	}

	return strconv.Itoa(safeSpaceCount), nil
}

// PartTwo solves day 18 part one.
func PartTwo(input string) (string, error) {
	currentRow := []rune(input)
	safeSpaceCount := countType(currentRow, safeSpace)

	for i := 0; i < 399999; i++ {
		currentRow = processRow(currentRow)
		safeSpaceCount += countType(currentRow, safeSpace)
	}

	return strconv.Itoa(safeSpaceCount), nil
}
