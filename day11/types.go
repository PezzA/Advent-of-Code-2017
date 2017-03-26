package day11

type reactor struct {
	microchip int
	generator int
}

type reactorList map[string]reactor

type facility struct {
	elevator int
	floors   int
	elements []string
	reactors reactorList
}

type solveState struct {
	depth          int
	state          facility
	previousStates []facility
	subStates      []solveState
}
