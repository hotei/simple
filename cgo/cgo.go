// cgo.go (c) David Rook 2010 - released under Simplified BSD 2-clause License
// currently failing

// +build ignore 

package main

// #include <stdlib.h>
// #include <stdio.h>
// #cgo LDFLAGS: -lmath

import "C"

import "fmt"

func Random() int {
	return int(C.random())
}

func Seed(i int) {
	C.srandom(C.uint(i))
}

func main() {
	Seed(0)
	fmt.Printf("Testing %d\n", Random())
}
