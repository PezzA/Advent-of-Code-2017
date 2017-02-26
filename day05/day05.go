package day05

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
)

func PuzzleInput() string {
	return `ffykfhsq`
}

func PartOne() (string, error) {
	hasher := md5.New()

	input := PuzzleInput()

	var password string

	for i := 0; len(password) < 8; i++ {
		hasher.Reset()
		hasher.Write([]byte(input + strconv.Itoa(i)))

		hash := hex.EncodeToString(hasher.Sum(nil))

		if hash[:5] == "00000" {
			password += hash[5:6]

		}
	}
	return password, nil
}

func PartTwo() (string, error) {
	hasher := md5.New()

	input := PuzzleInput()

	passMap := make(map[int]string)

	for i := 0; len(passMap) < 8; i++ {
		hasher.Reset()
		hasher.Write([]byte(input + strconv.Itoa(i)))

		hash := hex.EncodeToString(hasher.Sum(nil))

		if hash[:5] == "00000" {
			index, err := strconv.Atoi(hash[5:6])

			if err != nil {
				continue
			}

			val := hash[6:7]
			if index < 8 {
				if _, ok := passMap[index]; !ok {
					passMap[index] = val
				}
			}
		}
	}

	var password string

	for i := 0; i < 8; i++ {
		password += passMap[i]
	}
	return string(password), nil
}
