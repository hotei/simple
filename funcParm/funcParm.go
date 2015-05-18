// funcParm.go  (c) David Rook 2010 - released under Simplified BSD 2-clause License

package main

import (
	"fmt"
)

func addit(a, b int) int {
	//	fmt.Printf("addit\n")
	return a + b
}

func diffit(a, b int) int {
	//	fmt.Printf("diffit\n")
	return a - b
}

func funtest(x, y int, fn func(a, b int) int) int {
	z := fn(x, y)
	return z
}

func main() {
	fmt.Printf("Started\n")
	fmt.Printf("addit(%d) should be 10\n", addit(4, 6))
	fmt.Printf("diffit(%d) should be  4\n", diffit(10, 6))
	fmt.Printf("addit(%d) should be 12\n", funtest(3, 9, addit))
	fmt.Printf("diffit(%d) should be  3\n", funtest(12, 9, diffit))
	fmt.Printf("Finished\n")
}
