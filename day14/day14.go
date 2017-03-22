package day14

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"strings"
)

// PuzzleInput returns the puzzle input for day14
func PuzzleInput() string {
	return `ahsbgdzn`
}

// TestInput returns the test input for day14
func TestInput() string {
	return `abc`
}

var hashCache map[int]string

// This will store the result of each hash operation.  Each hash will be requested at least
// 1000 times so this will save an awful lot of processing
func getCachedHash(salt string, index int, keyStretchFactor int) string {
	if val, ok := hashCache[index]; ok {
		return val
	}

	hasher := md5.New()
	hasher.Write([]byte(salt + strconv.Itoa(index)))
	hash := hex.EncodeToString(hasher.Sum(nil))

	for i := 0; i < keyStretchFactor; i++ {
		hasher.Reset()
		hasher.Write([]byte(hash))
		hash = hex.EncodeToString(hasher.Sum(nil))

	}

	hashCache[index] = hash

	return hash
}

func hashSearch(salt string, startingIndex int, iterations int, comparason string, keyStretchFactor int) bool {
	var match = strings.Repeat(comparason[:1], 5)
	for i := 0; i <= iterations; i++ {
		hash := getCachedHash(salt, startingIndex+i, keyStretchFactor)
		for i := 0; i <= len(hash)-5; i++ {
			if hash[i:i+5] == match {
				return true
			}
		}
	}
	return false
}

func isKey(s string, index int, keyStretchFactor int) bool {
	hash := getCachedHash(s, index, keyStretchFactor)
	for i := 0; i <= len(hash)-3; i++ {
		chunk := hash[i : i+3]
		if chunk[:1] == chunk[1:2] && chunk[1:2] == chunk[2:3] {
			return hashSearch(s, index+1, 1000, chunk, keyStretchFactor)
		}
	}
	return false
}

func run(salt string, keyStretchFactor int) int {
	hashCache = make(map[int]string, 0)
	control := 0
	index := 0
	for {
		if isKey(salt, index, keyStretchFactor) {
			control++
			if control == 64 {
				break
			}
		}
		index++
	}
	return index
}

// PartTwo runs day14 part two.
func PartTwo(s string) (string, error) {
	return strconv.Itoa(run(s, 2016)), nil
}

// PartOne runs day14 part one.
func PartOne(s string) (string, error) {
	return strconv.Itoa(run(s, 0)), nil
}
