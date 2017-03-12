package day09

import (
	"regexp"
	"strconv"
)

var markerRegex, _ = regexp.Compile(`\((?P<Length>[0-9]*)x(?P<Reps>[0-9]*)\)`)

type section struct {
	data    string
	repeats int
}

func getRenderedLength(sx []section) int {

	length := 0

	for _, sec := range sx {
		length += len(sec.data) * sec.repeats
	}

	return length
}

func getLength(s string) int {
	markerPos := markerRegex.FindIndex([]byte(s))

	if markerPos == nil {
		return len(s)
	}

	match := markerRegex.FindAllStringSubmatch(s[markerPos[0]:markerPos[1]], -1)

	aft := s[markerPos[1]:]

	repeats, _ := strconv.Atoi(match[0][2])
	markerLen, _ := strconv.Atoi(match[0][1])

	return markerPos[0] + (repeats*getLength(aft[:markerLen]) + getLength(aft[markerLen:]))
}

func getNextSection(s string) ([]section, string) {
	sections := make([]section, 0)

	markerPos := markerRegex.FindIndex([]byte(s))

	if markerPos == nil {
		return append(sections, section{s, 1}), ""
	}

	fore := s[:markerPos[0]]
	marker := s[markerPos[0]:markerPos[1]]
	aft := s[markerPos[1]:]

	// add the prefix
	if len(fore) > 0 {
		sections = append(sections, section{fore, 1})
	}

	r2 := markerRegex.FindAllStringSubmatch(marker, -1)

	markerLen, _ := strconv.Atoi(r2[0][1])
	markerRepeats, _ := strconv.Atoi(r2[0][2])

	markerString := aft[:markerLen]

	sections = append(sections, section{markerString, markerRepeats})

	return sections, aft[markerLen:]
}

func getAllSections(s string) []section {
	sections := make([]section, 0)

	remain := s
	for remain != "" {
		var result []section
		result, remain = getNextSection(remain)
		sections = append(sections, result...)
	}

	return sections
}

// PartTwo runs day 09 part two
func PartTwo(input string) (string, error) {
	return strconv.Itoa(getLength(input)), nil
}

// PartOne runs day 09 part one
func PartOne(input string) (string, error) {
	return strconv.Itoa(getRenderedLength(getAllSections(input))), nil
}
