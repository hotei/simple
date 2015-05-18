// readInput.go (c) David Rook 2012 - released under Simplified BSD 2-clause License

// +build ignore

package main

import (
	"fmt"
	"os"
)

func FatalError(err error) {
	fmt.Printf("!Err-> %s\n", err)
	os.Exit(1)
}

// singleCharRead might return newline if no other input
//  otherwise first character if more than one on a line
// ^D will cause EOF to be printed and ? as the char returned - better choice is ...
func SingleCharRead() byte {
	var buf = []byte{0}
	_, err := os.Stdin.Read(buf)
	if err != nil {
		fmt.Printf("%v\n", err)
		return '?'
	}
	return buf[0]
}

func main() {
	for {
		fmt.Printf("<readInput>\n")
		c := SingleCharRead()
		if c == 'q' {
			fmt.Printf("Quitting\n")
			break
		}
		fmt.Printf("%c %02x\n", c, c)
	}
}
