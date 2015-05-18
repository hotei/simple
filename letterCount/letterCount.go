// letterCount.go (c) David Rook 2010 - released under Simplified BSD 2-clause License

package main

import (
	"fmt"
)

//
func LetterCount(buf []byte) [256]int {
	var counts [256]int

	for _, val := range buf {
		counts[val]++
	}
	return counts
}

var testBuf = []byte{'a', 'b', 'c', 'c', 'c'}

func main() {
	rv := LetterCount(testBuf)
	fmt.Printf("%v\n", rv)
	fmt.Printf("<done>\n")
}
