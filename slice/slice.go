// slice.go (c) David Rook 2010 - released under Simplified BSD 2-clause License

// test slice operations - working

package main

import (
	"fmt"
)

func main() {
	sl := [...]int{10, 20, 30, 40, 50}
	var xl []int
	var yl []int
	var zl []int

	fmt.Printf("simple slice\n")
	fmt.Printf("Slice sl %v \n", sl)
	// delete second element (value=30)
	ndx := 2

	xl = sl[0:ndx]
	fmt.Printf("Slice xl %v  (should be [10, 20]\n", xl)

	yl = sl[ndx+1:]
	fmt.Printf("Slice yl %v (should be [40, 50])\n", yl)

	zl = xl
	for _, val := range yl {
		zl = append(zl, val)
	}
	fmt.Printf("Slice zl %v (should be [10, 20, 40, 50])\n", zl)

}
