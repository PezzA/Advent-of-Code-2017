package day07

import (
	"strconv"
	"strings"
)

func reverseABA(s string) string {
	return s[1:2] + s[0:1] + s[1:2]
}

func getIPAddresses(s string) []string {
	return strings.Split(s, "\n")
}

func isABA(s string) (bool, []string) {
	abas := make([]string, 0)
	for i := 0; i < (len(s) - 2); i++ {
		check := s[i : i+3]
		if check[:1] == check[2:3] &&
			check[0:1] != check[1:2] {
			abas = append(abas, check)
		}
	}
	return len(abas) > 0, abas
}

func isABBA(s string) (bool, []string) {
	abbas := make([]string, 0)
	for i := 0; i < (len(s) - 3); i++ {
		check := s[i : i+4]
		if check[:1] == check[3:4] &&
			check[1:2] == check[2:3] &&
			check[0:1] != check[1:2] {
			abbas = append(abbas, check)
		}
	}
	return len(abbas) > 0, abbas
}

var tlsSeperator = func(c rune) bool {
	return c == '[' || c == ']'
}

func getHyperNetSequences(s string) []string {
	hnSeqs := make([]string, 0)
	for index, bit := range strings.FieldsFunc(s, tlsSeperator) {
		if index%2 == 1 {
			hnSeqs = append(hnSeqs, bit)
		}
	}
	return hnSeqs
}

func getSuperNetSequences(s string) []string {
	hnSeqs := make([]string, 0)
	for index, bit := range strings.FieldsFunc(s, tlsSeperator) {
		if index%2 == 0 {
			hnSeqs = append(hnSeqs, bit)
		}
	}
	return hnSeqs
}

func supportsSLS(s string) bool {
	snsAbaList := make([]string, 0)

	for _, bit := range getSuperNetSequences(s) {
		if aba, abaList := isABA(bit); aba {
			snsAbaList = append(snsAbaList, abaList...)
		}
	}

	hnsAbaList := make([]string, 0)
	for _, bit := range getHyperNetSequences(s) {
		if aba, abaList := isABA(bit); aba {
			hnsAbaList = append(hnsAbaList, abaList...)
		}
	}

	hasReverseABA := false

	for _, aba := range snsAbaList {
		for _, hnAba := range hnsAbaList {
			if hnAba == reverseABA(aba) {
				hasReverseABA = true
			}
		}
	}

	return hasReverseABA
}

func supportsTLS(s string) bool {
	containsAbbainSNS := false
	for _, bit := range getSuperNetSequences(s) {
		if isAbba, _ := isABBA(bit); isAbba {
			containsAbbainSNS = true
		}
	}

	containsAbbainHNS := false
	for _, bit := range getHyperNetSequences(s) {
		if isAbba, _ := isABBA(bit); isAbba {
			containsAbbainHNS = true
		}
	}
	return containsAbbainSNS && !containsAbbainHNS
}

// PartTwo runs PartTwo for day07
func PartTwo(input string) (string, error) {
	ips := getIPAddresses(input)

	valid := 0

	for _, ip := range ips {
		if supportsSLS(ip) {
			valid++
		}
	}

	return strconv.Itoa(valid), nil
}

// PartOne runs PartOne for day07
func PartOne(input string) (string, error) {
	ips := getIPAddresses(input)

	valid := 0

	for _, ip := range ips {
		if supportsTLS(ip) {
			valid++
		}
	}

	return strconv.Itoa(valid), nil
}
