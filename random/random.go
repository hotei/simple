// random.go (c) David Rook 2010 - released under Simplified BSD 2-clause License

package main

import (
	"fmt"
	"math/rand"
)

func main() {

	rand.Seed(0)

	sum := 0.0
	const N = 2000000
	for i := 0; i < N; i++ {
		//fmt.Printf("random number %d = %g\n", i+1, rand.Float64())
		sum += rand.Float64()
	}

	fmt.Printf("Average of %d random Float64 is %g, should be about .5000000000000 :-) \n", N, sum/N)

	var v float32 = 1.0 / 3.0
	fmt.Printf("v = %g (should be 0.3 repeating with 32bit precision = 8 digits) \n", v)

	var t float64 = 1.0 / 3.0
	fmt.Printf("t = %g (should be 0.3 repeating with 64bit precision = 16 digits) \n", t)
}
