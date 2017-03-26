package day14

import "testing"
import "log"

func Test_someStuff(t *testing.T) {
	log.Println(getCachedHash("abc", 0, 0))
	log.Println(getCachedHash("abc", 22551, 2016))
	log.Println(getCachedHash("abc", 22859, 2016))
}
