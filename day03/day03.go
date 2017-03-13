package day03

import (
	"strconv"
	"strings"
)

type triangle struct {
	side1 int
	side2 int
	side3 int
}

// PartOne executes part one of the puzzle
func PartOne(input string) (string, error) {
	validCounter := 0

	for _, t := range getTriangles(input) {
		if t.isValid() {
			validCounter++
		}
	}

	return strconv.Itoa(validCounter), nil
}

// PartTwo executes part two of the puzzle
func PartTwo(input string) (string, error) {
	validCounter := 0

	for _, t := range getColTriangles(input) {
		if t.isValid() {
			validCounter++
		}
	}

	return strconv.Itoa(validCounter), nil
}

func (t triangle) isValid() bool {
	return (t.side1+t.side2) > t.side3 && (t.side1+t.side3) > t.side2 && (t.side3+t.side2) > t.side1
}

func getColTriangles(input string) []triangle {
	triangles := make([]triangle, 0)

	var t1 triangle
	var t2 triangle
	var t3 triangle

	for i, line := range strings.Split(input, "\n") {
		if i%3 == 0 {
			t1 = lineToTriangle(line)
		}
		if i%3 == 1 {
			t2 = lineToTriangle(line)
		}
		if i%3 == 2 {
			t3 = lineToTriangle(line)

			triangles = append(triangles, triangle{t1.side1, t2.side1, t3.side1})
			triangles = append(triangles, triangle{t1.side2, t2.side2, t3.side2})
			triangles = append(triangles, triangle{t1.side3, t2.side3, t3.side3})
		}
	}

	return triangles
}

func lineToTriangle(line string) triangle {
	var bits = strings.Fields(strings.TrimSpace(line))

	side1, _ := strconv.Atoi(bits[0])
	side2, _ := strconv.Atoi(bits[1])
	side3, _ := strconv.Atoi(bits[2])

	sides := []int{side1, side2, side3}

	return triangle{sides[0], sides[1], sides[2]}
}

func getTriangles(input string) []triangle {
	triangles := make([]triangle, 0)

	for _, line := range strings.Split(input, "\n") {

		triangles = append(triangles, lineToTriangle(line))
	}

	return triangles
}
