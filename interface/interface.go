// interface.go (c) David Rook 2010 - released under Simplified BSD 2-clause License

/*
 * demonstrate functions that take multiple types as args
 * any argument type that satisfies the Showable interface def can be used as
 * argument to Show(arg).  Normally Show() would involve graphics but it's not
 * required to demonstrate the principle.
 */

package main

import "fmt"

type Showable interface {
	PrintSelf()
}

type rectangle struct {
	h int
	w int
}

type circle struct {
	R int
}

func (b *rectangle) PrintSelf() {
	fmt.Printf("I'm a rectangle height(%d) width(%d)\n", b.h, b.w)
}

func (b *circle) PrintSelf() {
	fmt.Printf("I'm a circle of radius(%d)\n", b.R)
}

func Show(x Showable) {
	x.PrintSelf()
}

func main() {
	fmt.Printf("<start SimpleInterface>\n")

	circ := new(circle)
	circ.R = 10
	rect := new(rectangle)
	rect.h = 4
	rect.w = 8
	circ.PrintSelf()
	rect.PrintSelf()
	// Show function can take more than one different type as argument
	Show(circ)
	Show(rect)

	fmt.Printf("<end SimpleInterface>\n")
}
