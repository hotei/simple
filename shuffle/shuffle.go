// shuffle.go (c) David Rook 2010 - released under Simplified BSD 2-clause License

// for each element, interchange with a random element
// it's possible to interchange with self and that's ok
package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var x = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := len(x) - 1; i >= 1; i-- {
		j := rand.Intn(i + 1)
		fmt.Printf("exchange [%d] and [%d]\n", i, j)
		x[i], x[j] = x[j], x[i]
		fmt.Printf("x=%v\n", x)
	}
}
