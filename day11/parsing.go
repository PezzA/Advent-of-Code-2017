package day11

import "strings"

func getStartingFacility(s string) facility {
	reactorList := make(reactorList)

	// find all the microchips in play
	for index, line := range strings.Split(s, "\n") {
		for _, word := range strings.Fields(line) {
			if strings.Index(word, "-") > 0 {
				element := strings.Split(word, "-")
				if _, ok := reactorList[element[0]]; !ok {
					reactorList[element[0]] = reactor{index, 0}
				}
			}
		}
	}

	// then go and find where the reactors are
	for index, line := range strings.Split(s, "\n") {
		for k, v := range reactorList {
			if strings.Index(line, k+" generator") > 0 {
				v.generator = index
				reactorList[k] = v
			}
		}
	}

	elements := make([]string, 0)
	for k := range reactorList {
		elements = append(elements, k)
	}

	return facility{0, 4, elements, reactorList}
}
