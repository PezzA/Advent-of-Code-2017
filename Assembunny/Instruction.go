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
	if i.isDualArg() {
		if i.action == jumpVal {
			return Instruction{copyVal, i.source, i.sourceVal, i.target, i.targetVal}
		}

		if i.action == jumpReg {
			return Instruction{copyReg, i.source, i.sourceVal, i.target, i.targetVal}
		}

		if i.sourceVal != 0 {
			return Instruction{jumpVal, i.source, i.sourceVal, i.target, i.targetVal}
		}

		return Instruction{jumpReg, i.source, i.sourceVal, i.target, i.targetVal}
	}

	if i.action == dec {
		return Instruction{inc, i.source, i.sourceVal, i.target, i.targetVal}
	}

	if i.action == inc {
		return Instruction{dec, i.source, i.sourceVal, i.target, i.targetVal}
	}

	return Instruction{inc, i.source, i.sourceVal, i.target, i.targetVal}
}

func (i Instruction) isDualArg() bool {
	return i.targetVal != 0 || i.target != ""
}
