// readRandom.go (c) David Rook 2010 - released under Simplified BSD 2-clause License

package main

import (
	"fmt"
	"os"
)

func FatalError(err error) {
	fmt.Printf("!Err-> %s\n", err)
	os.Exit(1)
}

func main() {
	var buf = make([]byte, 8)

	input, err := os.Open("/dev/random")
	if err != nil {
		FatalError(err)
	}
	n, err := input.Read(buf)
	if (err != nil) || (n != 8) {
		FatalError(err)
	}
	fmt.Printf("%v\n", buf)
	fmt.Printf("</readRandom>\n")
}
