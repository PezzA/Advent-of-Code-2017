package day07

import (
	"strings"
)

func isABBA(s string) bool {
	for i := 0; i < (len(s) - 3); i++ {
		check := s[i : i+4]
		if check[:1] == check[3:4] &&
			check[1:2] == check[2:3] {
			return true
		}
	}
	return false
}

var tlsSeperator = func(c rune) bool {
	return c == '[' || c == ']'
}

func getHyperNetSequences(s string) []string {
	hnSeqs := make([]string, 0)
	for index, bit := range strings.FieldsFunc(s, tlsSeperator) {
		if index%2 == 0 {
			hnSeqs = append(hnSeqs, bit)
		}
	}
	return hnSeqs
}

func getIPSequences(s string) []string {
	hnSeqs := make([]string, 0)
	for index, bit := range strings.FieldsFunc(s, tlsSeperator) {
		if index%2 == 1 {
			hnSeqs = append(hnSeqs, bit)
		}
	}
	return hnSeqs
}

func supportsTLS(s string) bool {

	for _, bit := range getIPSequences(s) {
		if !isABBA(bit) {
			return false
		}
	}

	for _, bit := range getHyperNetSequences(s) {
		if isABBA(bit) {
			return false
		}
	}
	return true
}
