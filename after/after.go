// after.go (c) David Rook 2010 - released under Simplified BSD 2-clause License

// do something inside loop, break out of loop after specified time
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("Tick\n")

	n := 0
	ch := time.After(10 * time.Second)
L1:
	for {
		fmt.Printf("Tock %d\n", n)
		n++
		select {
		case <-ch:
			fmt.Printf("Timed out\n")
			break L1 // label required
		default: // nothing
		}
		time.Sleep(3 * time.Second)
	}
	fmt.Printf("Chime\n")
}
