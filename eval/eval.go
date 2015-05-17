// eval.go (c) David Rook 2010 - released under Simplified BSD 2-clause License

/*
 worked ok with function as variable
*/

package main

import (
	"fmt"
	"os"
)

func Pause() {
	fmt.Printf("Hit a key to continue\n")
	var keybuf [1]byte
	n, err := os.Stdin.Read(keybuf[0:1])
	if err != nil {
		fmt.Printf("got error %v and %d chars\n", err, n)
	}
}

func tester(d int) {
	fmt.Printf("Test from func tester() %d\n", d)
}

func main() {
	var y func(int)
	fmt.Printf("Test from main() \n")
	x := tester
	tester(2)
	x(3)
	y = tester
	y(4)
}
