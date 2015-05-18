// ticker.go (c) David Rook 2010 - released under Simplified BSD 2-clause License

package main

import (
	"fmt"
	"time"
)

func statusUpdate() string {
	return fmt.Sprintf("Status\n")
}

func main() {
	fmt.Printf("Tick\n")
	c := time.Tick(3 * time.Second)
	for now := range c {
		fmt.Printf("%v %s\n", now, statusUpdate())
	}
	fmt.Printf("Tock\n")
}
