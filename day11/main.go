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

func (f facility) isValid(elevatorContents []component, floor int) (bool, facility) {
	chips := make([]component, 0)

	allComponents := append(f.getItemsOnFloor(floor), elevatorContents...)
	for _, component := range allComponents {
		if component.isMicrochip {
			chips = append(chips, component)
		}
	}

	for _, chip := range chips {
		// for each chip the move is invalid if the floor
		//	- does not the chips generator AND contains another generator
		containsOwnGenerator := false
		containsOtherGenerator := false

		for _, component := range allComponents {
			if !component.isMicrochip {
				if component.element == chip.element {
					containsOwnGenerator = true
				} else {
					containsOtherGenerator = true
				}
			}
			component.floor++
		}

		if !containsOwnGenerator && containsOtherGenerator {
			return false, f
		}
	}

	updatedComponentList := make([]component, 0)
	for _, ax := range f.components {
		for _, bx := range elevatorContents {
			if ax.element == bx.element &&
				ax.isMicrochip == bx.isMicrochip {
				updatedComponentList = append(updatedComponentList, component{floor, ax.element, ax.isMicrochip})
			} else {
				updatedComponentList = append(updatedComponentList, ax)
			}
		}
	}
	return true, facility{floor, updatedComponentList}
}

//get all the items on the specified floor
func (f facility) getItemsOnFloor(floor int) []component {
	components := make([]component, 0)

	for _, comp := range f.components {
		if comp.floor == floor {
			components = append(components, comp)
		}
	}
	return components
}

func (f facility) canGoDown() bool {
	if f.elevator > 0 {
		return true
	}
	return false
}

func (f facility) canGoUp() bool {
	if f.elevator < 2 {
		return true
	}
	return false
}

func solve(inputState facility) solveState {
	state := solveState{inputState, make([]solveState, 0)}
	moves := elevatorCombinations(inputState.getItemsOnFloor(inputState.elevator))

	// for each combination of items we can put in the elevator, move them up and down (if able too)
	for _, move := range moves {
		if inputState.canGoUp() {
			if valid, newFacility := inputState.isValid(move, inputState.elevator+1); valid {
				state.subStates = append(state.subStates, solveState{newFacility, make([]solveState, 0)})
			}
		}
		if inputState.canGoDown() {
			if valid, newFacility := inputState.isValid(move, inputState.elevator-1); valid {
				state.subStates = append(state.subStates, solveState{newFacility, make([]solveState, 0)})
			}
		}
	}
	return state
}

// Given a list of items, returns a list of
// all the possibly combinations of items
// that could be put into the elevator.
func elevatorCombinations(componentList []component) [][]component {
	parts := make([][]component, 0)
	maxCard := int(math.Pow(2, float64(len(componentList))))

	for i := 0; i < maxCard; i++ {
		takeParts := make([]component, 0)
		count := strings.Count(strconv.FormatInt(int64(i), 2), "1")

		if count == 1 || count == 2 {
			for j := 0; j < len(componentList); j++ {
				if i&(1<<uint(j)) != 0 {
					takeParts = append(takeParts, componentList[j])
				}
			}
		}
		if len(takeParts) > 0 {
			parts = append(parts, takeParts)
		}
	}
	return parts
}
