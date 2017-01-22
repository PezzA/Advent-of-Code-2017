package main

import (
	"errors"
	"strconv"
	"strings"
)

type instruction struct {
	direction string
	steps     int
}

type point struct {
	x int
	y int
}

type position struct {
	p      point
	facing int
}

func intabs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func getlength(p point) int {
	return intabs(p.x) + intabs(p.y)
}

// PartOne solves the first part of day one
func PartOne(inputData string) int {
	var currentPosition = position{point{0, 0}, 0}

	for _, ins := range parseInstructions(inputData) {
		currentPosition, _ = move(currentPosition, ins)
	}

	return getlength(currentPosition.p)
}

func containsPoint(points []point, point point) bool {
	for _, testPoint := range points {
		if (testPoint.x == point.x) && (testPoint.y == point.y) {
			return true
		}
	}
	return false
}

// PartTwo solves the second part of day one
func PartTwo(inputData string) (int, error) {
	var currentPosition = position{point{0, 0}, 0}

	points := make([]point, 0)

	for _, ins := range parseInstructions(inputData) {
		newPoints := make([]point, 0)

		currentPosition, newPoints = move(currentPosition, ins)

		for _, testPoint := range newPoints {
			if containsPoint(points, testPoint) {
				return getlength(testPoint), nil
			}
		}

		points = append(points, newPoints...)
	}

	return 0, errors.New("no intersections found")
}

func move(pos position, ins instruction) (position, []point) {
	var newOrientation = orientate(ins.direction, pos.facing)

	points := make([]point, 0)

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

func translate(orientation int, shift int, startPoint point) point {
	switch orientation {
	case 0:
		return point{startPoint.x, startPoint.y + shift}
	case 90:
		return point{startPoint.x + shift, startPoint.y}
	case 180:
		return point{startPoint.x, startPoint.y - shift}
	case 270:
		return point{startPoint.x - shift, startPoint.y}
	}

	return point{0, 0}
}
