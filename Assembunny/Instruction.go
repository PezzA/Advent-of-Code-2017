package Assembunny

type instructionType uint8

const (
	copyVal instructionType = iota
	copyReg instructionType = iota
	inc     instructionType = iota
	dec     instructionType = iota
	jumpReg instructionType = iota
	jumpVal instructionType = iota
	toggle  instructionType = iota
)

// Instruction represents a single Assembunny instruction.
type Instruction struct {
	action    instructionType
	source    string
	sourceVal int
	target    string
	targetVal int
}

func (i Instruction) toggle() Instruction {
	return i
}

func (i Instruction) isDualArg() bool {
	return i.targetVal != 0 || i.target != ""
}
