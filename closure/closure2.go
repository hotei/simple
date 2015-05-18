// closure.go (c) David Rook 2010 - released under Simplified BSD 2-clause License

// +build ignore

package main

import (
	"fmt"
)

func makecounterA(c int) func() int { return func() int { c++; return c } }

func makecounterB(c int) func(i int) int { return func(i int) int { c += i; return c } }

func main() {

	f := makecounterA(0)
	g := makecounterA(30)
	fmt.Printf("f(%d) g(%d) expect f(1) g(31)\n", f(), g())
	fmt.Printf("f(%d) expect f(2)\n", f())
	fmt.Printf("f(%d) g(%d) expect f(3) g(32)\n", f(), g())

	x := makecounterB(0)
	fmt.Printf("x(%d) expect x(1) \n", x(1))
	fmt.Printf("x(%d) expect x(2) \n", x(1))
	fmt.Printf("x(%d) expect x(0) \n", x(-x(0)))
	//fmt.Printf("f(%d) expect f(34)\n", f(33))
	//fmt.Printf("f(%d) g(%d) expect f(34) g(31)\n", f(0), g(0))

}
