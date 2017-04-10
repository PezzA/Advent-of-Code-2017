package day17

import (
	"crypto/md5"
	"encoding/hex"
)

// starts at zero
var gridSize = 3

type point struct {
	x int
	y int
}

// PartOne solves day 17 part one.
func PartOne(input string) (string, error) {
	return "", nil
}

// PartTwo solves day 17 part two.
func PartTwo(input string) (string, error) {
	return "", nil
}

func getHash(input string) string {
	hasher := md5.New()
	hasher.Write([]byte(input))
	hash := hex.EncodeToString(hasher.Sum(nil))
	return hash
}

func isOpen(input rune) bool {
	return input >= 98 && input <= 102
}

func getDoors(input string, pos point) []bool {
	hash := []rune(getHash(input)[:4])

	up := pos.x != 0 && isOpen(hash[:1][0])
	down := pos.x != gridSize && isOpen(hash[1:2][0])
	left := pos.y != 0 && isOpen(hash[2:3][0])
	right := pos.y != gridSize && isOpen(hash[3:4][0])

	return []bool{up, down, left, right}
}
