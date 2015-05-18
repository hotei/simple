// memoize.go (c) David Rook 2010 - released under Simplified BSD 2-clause License

// demonstrates general method for creating memoizing functions

package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

type BinaryPointFn func(Point, Point) float64

var Dist = func(p1, p2 Point) float64 {
	fmt.Printf("calculating new value %v %v ...", p1, p2)
	return math.Sqrt(math.Pow(p2.X-p1.Y, 2.0) + math.Pow(p2.Y-p1.Y, 2.0))
}

// Memoize must be tailored for the function being replaced
func Memoize(fn BinaryPointFn) BinaryPointFn {
	type T struct {
		p1 Point
		p2 Point
	}
	history := make(map[T]float64)
	return func(p1, p2 Point) float64 {
		if res, ok := history[T{p1, p2}]; ok {
			fmt.Printf("reading from history %v %v ...", p1, p2)
			return res
		}
		// same result if order of points is switched so check that too
		if res, ok := history[T{p2, p1}]; ok {
			fmt.Printf("reading from history %v %v ...", p1, p2)
			return res
		}
		// not in history so do it the hard way
		res := fn(p1, p2)
		// and save the result for later
		history[T{p1, p2}] = res
		return res
	}
}

func main() {
	p1 := Point{3.1, 10.3}
	p2 := Point{-4.7, 28.9}
	p3 := Point{2.0, 20.}
	Dist = Memoize(Dist)
	fmt.Println(Dist(p1, p2))
	fmt.Println(Dist(p1, p3))
	fmt.Println(Dist(p2, p1))
	fmt.Println(Dist(p1, p2))
	fmt.Println(Dist(p2, p1))
}
