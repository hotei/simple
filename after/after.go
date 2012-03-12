// after.go
// (c) David Rook 2012 - released under Simplified BSD Licens - see file "after.md"

// do something inside loop, break out of loop after specified time
package main

import (
	"fmt"
	"time"
)

var (
	license = "after.go pkg (c) 2012 David Rook released under Simplified BSD License"
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
