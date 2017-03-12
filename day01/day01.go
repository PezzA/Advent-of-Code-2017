package day01

import (
	"errors"
	"strconv"
	"strings"
)

// PuzzleInput returns the main puzzle input for day01
func PuzzleInput() string {
	return "R1, R1, R3, R1, R1, L2, R5, L2, R5, R1, R4, L2, R3, L3, R4, L5, R4, R4, R1, L5, L4, R5, R3, L1, R4, R3, L2, L1, R3, L4, R3, L2, R5, R190, R3, R5, L5, L1, R54, L3, L4, L1, R4, R1, R3, L1, L1, R2, L2, R2, R5, L3, R4, R76, L3, R4, R191, R5, R5, L5, L4, L5, L3, R1, R3, R2, L2, L2, L4, L5, L4, R5, R4, R4, R2, R3, R4, L3, L2, R5, R3, L2, L1, R2, L3, R2, L1, L1, R1, L3, R5, L5, L1, L2, R5, R3, L3, R3, R5, R2, R5, R5, L5, L5, R2, L3, L5, L2, L1, R2, R2, L2, R2, L3, L2, R3, L5, R4, L4, L5, R3, L4, R1, R3, R2, R4, L2, L3, R2, L5, R5, R4, L2, R4, L1, L3, L1, L3, R1, R2, R1, L5, R5, R3, L3, L3, L2, R4, R2, L5, L1, L1, L5, L4, L1, L1, R1"
}

// Point represents a basic 2d co-ordinate
type Point struct {
	x int
	y int
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

type instruction struct {
	direction string
	steps     int
}

type position struct {
	p      Point
	facing int
}

func getlength(p Point) int {
	return abs(p.x) + abs(p.y)
}

// PartOne returns the solution for day01 part one
func PartOne(input string) (string, error) {
	return solveRawLocation(input)
}

func solveRawLocation(inputData string) (string, error) {
	var currentPosition = position{Point{0, 0}, 0}

	for _, ins := range parseInstructions(inputData) {
		currentPosition, _ = move(currentPosition, ins)
	}

	return strconv.Itoa(getlength(currentPosition.p)), nil
}

func containsPoint(points []Point, point Point) bool {
	for _, testPoint := range points {
		if (testPoint.x == point.x) && (testPoint.y == point.y) {
			return true
		}
	}
	return false
}

// PartTwo returns the solution for day01 part two
func PartTwo(input string) (string, error) {
	return solveActualLocation(input)
}

func solveActualLocation(inputData string) (string, error) {
	var currentPosition = position{Point{0, 0}, 0}

	points := make([]Point, 0)

	for _, ins := range parseInstructions(inputData) {
		newPoints := make([]Point, 0)

		currentPosition, newPoints = move(currentPosition, ins)

		for _, testPoint := range newPoints {
			if containsPoint(points, testPoint) {
				return strconv.Itoa(getlength(testPoint)), nil
			}
		}

		points = append(points, newPoints...)
	}

	return "-", errors.New("no intersections found")
}

func move(pos position, ins instruction) (position, []Point) {
	var newOrientation = orientate(ins.direction, pos.facing)

	points := make([]Point, 0)

	newPoint := pos.p

	for i := 0; i < ins.steps; i++ {
		newPoint = translate(newOrientation, 1, newPoint)
		points = append(points, newPoint)
	}

	return position{newPoint, newOrientation}, points
}

func parseInstructions(input string) []instruction {
	instructions := make([]instruction, 0)

	for _, c := range strings.Split(input, ",") {
		var input = strings.TrimSpace(c)

		step, _ := strconv.Atoi(input[1:len(input)])

		instructions = append(instructions, instruction{direction: input[0:1], steps: step})
	}

	return instructions
}

func orientate(direction string, currentOrientation int) int {
	switch direction {
	case "L":
		if currentOrientation == 0 {
			return 270
		}
		return currentOrientation - 90
	case "R":
		if currentOrientation == 270 {
			return 0
		}
		return currentOrientation + 90
	default:
		return 0
	}
}

func translate(orientation int, shift int, startPoint Point) Point {
	switch orientation {
	case 0:
		return Point{startPoint.x, startPoint.y + shift}
	case 90:
		return Point{startPoint.x + shift, startPoint.y}
	case 180:
		return Point{startPoint.x, startPoint.y - shift}
	case 270:
		return Point{startPoint.x - shift, startPoint.y}
	}

	return Point{0, 0}
}
