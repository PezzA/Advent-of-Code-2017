package day11

import (
	"fmt"
	"strconv"
	"strings"
)

func (f facility) getLine(floor int) string {
	line := ""

	line += "F" + strconv.Itoa(floor+1) + " "

	if f.elevator == floor {
		line += "E  "
	} else {
		line += ".  "
	}

	for _, element := range f.elements {
		initial := strings.ToUpper(element[:1])
		if f.reactors[element].generator == floor {
			line += initial + "G  "
		} else {
			line += ".   "
		}
		if f.reactors[element].microchip == floor {
			line += initial + "M  "
		} else {
			line += ".   "
		}
	}
	return line
}

func (f facility) drawFacility() {
	fmt.Println("----")
	for i := f.floors - 1; i >= 0; i-- {
		fmt.Println(f.getLine(i))
	}
}
