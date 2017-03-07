package day07

import "testing"
import "log"

func Test_AbbaDetection(t *testing.T) {
	if !isABBA("fgbabbaer") {
		t.Errorf("should have detected abba")
	}

	if !isABBA("ertyuuy") {
		t.Errorf("should have detected abba")
	}

	if isABBA("dfgdfgertertdfg") {
		t.Errorf("should not have detected abba")
	}

	log.Println(getHyperNetSequences("wysextplwqpvipxdv[srzvtwbfzqtspxnethm]syqbzgtboxxzpwr[kljvjjkjyojzrstfgrw]obdhcczonzvbfby[svotajtpttohxsh]cooktbyumlpxostt"))
	log.Println(getIPSequences("wysextplwqpvipxdv[srzvtwbfzqtspxnethm]syqbzgtboxxzpwr[kljvjjkjyojzrstfgrw]obdhcczonzvbfby[svotajtpttohxsh]cooktbyumlpxostt"))

	log.Println(supportsTLS("abba[mnop]qrst"))
	log.Println(supportsTLS("abcd[bddb]xyyx"))
	log.Println(supportsTLS("aaaa[qwer]tyui"))
	log.Println(supportsTLS("ioxxoj[asdfgh]zxcvbn"))
}
