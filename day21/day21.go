package day21

import "strings"

type password []rune

func getRotatedPos(length int, shift int, position int) int {
	actualShift := shift % length
	return (position + actualShift) % length
}

func (p password) toString() string {
	return string(p)
}

func (p password) rotate(x int) {
	temp := make(password, len(p))
	copy(temp, p)
	for index := range p {
		p[index] = temp[getRotatedPos(len(p), x, index)]
	}
}

func (p password) swapPos(x int, y int) {
	temp := p[x]
	p[x] = p[y]
	p[y] = temp
}

func (p password) swapChar(x rune, y rune) {
	posX := strings.IndexRune(string(p), rune(x))
	posY := strings.IndexRune(string(p), rune(y))
	p.swapPos(posX, posY)
}

// PartOne solves day 21 part one.
func PartOne(input string) (string, error) {
	return "", nil
}

// PartTwo solves day 21 part one.
func PartTwo(input string) (string, error) {
	return "", nil
}
