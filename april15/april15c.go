// april15c.go (c) David Rook 2010 - released under Simplified BSD 2-clause License

// NOTE: the executable is april15, not april15c

// NOTE !!! requires over 8 GB RAM to run.

// Works fine with 16 GB.  The key point is to use slices and allocate with make.
// do NOT use var taxpayerW#[MaxSSN]uint8 because that WILL fail after W2 no
// matter how much RAM you have.

package main

import (
	"fmt"
	"time"
)

const (
	MaxSSN = 1000 * 1000 * 1000 // USA SSN has 1e+9 entries possible
	DELAY  = 20
)

var (
	taxpayerW1 []uint8
	taxpayerW2 []uint8
	taxpayerW3 []uint8
	taxpayerW4 []uint8
	taxpayerW5 []uint8
	taxpayerW6 []uint8
	taxpayerW7 []uint8
	taxpayerW8 []uint8
)

func init() {
	fmt.Printf("init W1\n")
	taxpayerW1 = make([]uint8, MaxSSN)
	fmt.Printf("init W2\n")
	taxpayerW2 = make([]uint8, MaxSSN)
	fmt.Printf("init W3\n")
	taxpayerW3 = make([]uint8, MaxSSN)
	fmt.Printf("init W4\n")
	taxpayerW4 = make([]uint8, MaxSSN)
	fmt.Printf("init W5\n")
	taxpayerW5 = make([]uint8, MaxSSN)
	fmt.Printf("init W6\n")
	taxpayerW6 = make([]uint8, MaxSSN)
	fmt.Printf("init W7\n")
	taxpayerW7 = make([]uint8, MaxSSN)
	fmt.Printf("init W8\n")
	taxpayerW8 = make([]uint8, MaxSSN)

}
func main() {
	fmt.Printf("april15.go initializing each array with %d elements\n", MaxSSN)
	for i := 0; i < MaxSSN; i++ {
		taxpayerW1[i] = 1
		taxpayerW2[i] = 2
		taxpayerW3[i] = 3
		taxpayerW4[i] = 4
		taxpayerW5[i] = 5
		taxpayerW6[i] = 6
		taxpayerW7[i] = 7
		taxpayerW8[i] = 8
	}

	fmt.Printf("april15.go finished\n")
	time.Sleep(DELAY * time.Second)
}
