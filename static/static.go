// static.go (c) David Rook 2010 - released under Simplified BSD 2-clause License

package main

import (
	"fmt"
)

// counter returns starting value plus one
func counter(c int) func() int { return func() int { c++; return c } }

// static returns starting value plus increment
func static(c int) func(i int) int { return func(i int) int { c += i; return c } }

func main() {
	f := counter(2)
	g := counter(0)
	inc := static(15)

	fmt.Printf("inc(%d) \n", inc(4))
	fmt.Printf("inc(%d) \n", inc(-6))
	fmt.Printf("f(%d) g(%d)\n", f(), g())
	fmt.Printf("f(%d)\n", f())
	fmt.Printf("f(%d) g(%d)\n", f(), g())
}
