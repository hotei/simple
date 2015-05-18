// struct.go (c) David Rook 2010 - released under Simplified BSD 2-clause License

/*
 * example of struct initialization
 */

package main

import (
	"fmt"
	"math"
)

type Point2Di struct{ x, y int }

type Point2D struct{ x, y float64 }

func (p *Point2D) Dist2() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}

func dist_test() {
	pt := Point2D{10.0, 10.0}
	fmt.Printf("distance from origin to pt = %g (should be 14.142135623730951)\n", pt.Dist2())
}

func twod_int() {
	t := new(Point2Di)
	t.x = 10
	t.y = 20
	fmt.Printf("here [%d %d]\n", t.x, t.y)
}

func twod_real() {

	type Point2Di struct{ x, y float64 } // local struct with same name overrides global

	t := new(Point2Di)
	t.x = 10.5
	t.y = 20.5
	fmt.Printf("here [%g %g]\n", t.x, t.y)
}

func ary_test_one() { // array of integers

	const N = 25
	var ary [N]int

	for i := 0; i < N; i++ {
		ary[i] = i
	}

	for i := 0; i < N; i++ {
		fmt.Printf("ary[%d] = %d\n", i, ary[i])
	}
}

func ary_test_two() { // array of points

	const N = 25
	var ary [N]Point2Di

	for i := 0; i < N; i++ {
		ary[i].x = i
		ary[i].y = i * 2
	}

	for i := 0; i < N; i++ {
		fmt.Printf("ary[%d] = [%d %d]\n", i, ary[i].x, ary[i].y)
	}
}

func main() {
	twod_int()
	twod_real()
	ary_test_one()
	ary_test_two()
	dist_test()
}
