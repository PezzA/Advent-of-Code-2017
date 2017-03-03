package day06

import (
	"sort"
	"strings"

	mapsorter "github.com/pezza/advent-of-go/mapSorter"
)

type charMap map[rune]int

func parseLines(input string) []string {
	return strings.Split(input, "\n")
}

func getCharMaps([]string) []charMap {
	arrs := make([]charMap, 8)

	for index := range arrs {
		arrs[index] = make(charMap)
	}

	for _, entry := range parseLines(PuzzleInput()) {
		for i, char := range entry {
			if _, ok := arrs[i][char]; !ok {
				arrs[i][char] = 0
			}
			arrs[i][char]++
		}
	}

	return arrs
}

// PartOne runs partone!
func PartOne() (string, error) {
	charmaps := getCharMaps(parseLines(PuzzleInput()))

	var code string

	for _, val := range charmaps {
		sortedList := mapsorter.MapToList(val)
		sort.Sort(sort.Reverse(sortedList))
		code += string(sortedList[0].Key)
	}

	return code, nil
}

// PartTwo runs parttwo!
func PartTwo() (string, error) {
	charmaps := getCharMaps(parseLines(PuzzleInput()))

	var code string

	for _, val := range charmaps {
		sortedList := mapsorter.MapToList(val)
		sort.Sort(sortedList)
		code += string(sortedList[0].Key)
	}

	return code, nil
}
