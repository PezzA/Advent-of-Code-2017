package day04

type Pair struct {
	Key   rune
	Value int
}
type PairList []Pair

func (p PairList) Less(i, j int) bool {

	if p[i].Value == p[j].Value {
		return p[i].Key > p[j].Key
	}

	return p[i].Value < p[j].Value
}

func (p PairList) Len() int { return len(p) }

func (p PairList) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
