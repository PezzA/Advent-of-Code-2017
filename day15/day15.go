package day15

import "strconv"

type disc struct {
	positions   int
	startingPos int
}

func (d disc) GetTimeLaspedPosition(time int) int {
	return (d.startingPos + time) % d.positions
}

func dropBall(startingTime int, discs []disc) bool {
	for _, disc := range discs {
		startingTime++
		if disc.GetTimeLaspedPosition(startingTime) != 0 {
			return false
		}
	}
	return true
}

// PartTwo runs day15 part two.
func PartTwo(s string) (string, error) {
	discs := make([]disc, 7)

	discs[0] = disc{7, 0}
	discs[1] = disc{13, 0}
	discs[2] = disc{3, 2}
	discs[3] = disc{5, 2}
	discs[4] = disc{17, 0}
	discs[5] = disc{19, 7}
	discs[6] = disc{11, 0}

	time := 0
	for {
		if dropBall(time, discs) {
			break
		}
		time++
	}
	return strconv.Itoa(time), nil
}

// PartOne runs day15 part one.
func PartOne(s string) (string, error) {
	discs := make([]disc, 6)

	discs[0] = disc{7, 0}
	discs[1] = disc{13, 0}
	discs[2] = disc{3, 2}
	discs[3] = disc{5, 2}
	discs[4] = disc{17, 0}
	discs[5] = disc{19, 7}

	time := 0
	for {
		if dropBall(time, discs) {
			break
		}
		time++
	}
	return strconv.Itoa(time), nil
}
