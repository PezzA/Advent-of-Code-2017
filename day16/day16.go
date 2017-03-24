package day16

import (
	"log"
)

func expand(a []bool) []bool {
	b := make([]bool, 0)

	for i := len(a) - 1; i >= 0; i-- {
		if a[i] {
			b = append(b, false)
		} else {
			b = append(b, true)
		}
	}

	return append(append(a, false), b...)
}

func fill(input []bool, length int) []bool {
	retval := expand(input)
	for len(retval) < length {
		retval = expand(retval)
		log.Println("filling", len(retval))
	}

	return retval[:length]
}

func reduce(input []bool) []bool {
	retval := make([]bool, 0)

	for i := 0; i <= len(input)-1; i += 2 {
		if input[i] == input[i+1] {
			retval = append(retval, true)
		} else {
			retval = append(retval, false)
		}
	}
	return retval
}

func checkSum(input []bool) []bool {
	retval := reduce(input)

	for (len(retval) % 2) == 0 {
		retval = reduce(retval)
		log.Println("reducing", len(retval))
	}
	return retval
}

func boolToString(input []bool) string {
	retval := ""
	for i := 0; i <= len(input)-1; i++ {
		if input[i] {
			retval += "1"
		} else {
			retval += "0"
		}
	}

	return retval
}

func stringToBool(input string) []bool {
	bools := make([]bool, 0)

	for i := 0; i <= len(input)-1; i++ {
		if input[i] == 49 {
			bools = append(bools, true)
		} else {
			bools = append(bools, false)
		}
	}
	return bools
}

// PartTwo runs day16 part two.
func PartTwo(s string) (string, error) {
	data := fill(stringToBool(s), 35651584)
	return boolToString(checkSum(data)), nil
}

// PartOne runs day16 part one.
func PartOne(s string) (string, error) {
	data := fill(stringToBool(s), 272)
	return boolToString(checkSum(data)), nil
}
