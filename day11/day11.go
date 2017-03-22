package day11

import (
	"fmt"
	"strconv"
	"strings"

	ct "github.com/daviddengcn/go-colortext"
)

type reactor struct {
	microchip int
	generator int
	discover  int
}

type reactorList map[string]reactor

func (r reactorList) Len() int {
	return len(r)
}

type facility struct {
	elevator int
	floors   int
	reactors reactorList
}

func getStartingFacility(s string) facility {
	reactorList := make(reactorList)

	// find all the microchips in play
	for index, line := range strings.Split(s, "\n") {
		for _, word := range strings.Fields(line) {
			if strings.Index(word, "-") > 0 {
				element := strings.Split(word, "-")
				if _, ok := reactorList[element[0]]; !ok {
					reactorList[element[0]] = reactor{index, 0, len(reactorList)}
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

	return facility{0, 4, reactorList}
}

func drawFacility(f facility) {
	floors := make([]string, f.floors)

	for i := 0; i < f.floors; i++ {
		floors[i] = "F" + strconv.Itoa(i+1) + " " + strings.Repeat(".--", len(f.reactors))
	}

	for index := range floors {
		ct.Foreground(ct.Green, f.elevator == index)
		fmt.Print("F" + strconv.Itoa(index+1))

		fmt.Print("\n")

	}
	ct.Foreground(ct.White, true)
}

// PartOne returns the solution to day11 part one
func PartOne(input string) (string, error) {
	drawFacility(getStartingFacility(input))
	return "", nil
}

// PartTwo returns the solution to day11 part two
func PartTwo(input string) (string, error) {
	return "", nil
}
