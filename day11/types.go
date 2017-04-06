package day11

import (
	"strings"
)

type facility struct {
	elevator   int
	floors     int
	components []component
}

type component struct {
	floor       int
	element     string
	isMicrochip bool
}

func (c component) getInitial() string {
	return strings.ToUpper(c.element[:1])
}

func (c component) shortName() string {
	if c.isMicrochip {
		return c.getInitial() + "M"
	}

	return c.getInitial() + "G"
}
