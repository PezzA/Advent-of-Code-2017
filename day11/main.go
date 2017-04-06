package day11

// PartOne returns the solution to day11 part one
func PartOne(input string) (string, error) {
	return "", nil
}

// PartTwo returns the solution to day11 part two
func PartTwo(input string) (string, error) {
	return "", nil
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

/*
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
*/
