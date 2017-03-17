package day14

import "testing"
import "log"

func Test_someStuff(t *testing.T) {
	log.Println(getHash("abc", 0, 0))
	log.Println(getHash("abc", 22551, 2016))
	log.Println(getHash("abc", 22859, 2016))
}
