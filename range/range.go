// range.go (c) David Rook 2010 - released under Simplified BSD 2-clause License

/*
 * demo program of the flow control statement 'for'
 * also some array initializer examples
 */

package main

import "fmt"

func summit() {

	sum := 0
	for i := 1; i <= 9; i++ {
		sum += i
	}

	fmt.Printf("sum of numbers 1 to 9 = %d\n", sum)
}

// Note use of 'empty' return parameter (underscore)
func sumi(a []int) (s int) {
	for _, v := range a {
		s += v
	}
	return
}

func sumf(a []float32) (s float32) {
	for _, v := range a {
		s += v
	}
	return
}

func main() {
	array := [...]float32{7.0, 8.5, 9.1}
	numbers := [...]int{7, 8, 9}
	var moreNumbers = [...]int{1, 2, 3}

	fmt.Printf("listing index and value for each element in array []float \n")
	fmt.Printf("There are %d elements in the array \n", len(array))
	for index, value := range array {
		fmt.Printf("index(%v) has value(%v) ie. array[%d] = %g\n", index, value, index, value)
	}
	fmt.Println()
	var fsum float32 = sumf(array[0:])
	fmt.Printf("fsum is %g (should be 24.6)\n", fsum)
	fmt.Println()

	var isum int = sumi(numbers[0:])
	fmt.Printf("isum is %d (should be 24)\n", isum)
	fmt.Println()

	isum = sumi(moreNumbers[0:])
	fmt.Printf("isum is %d (should be 6)\n", isum)
	fmt.Println()

	summit()
	fmt.Printf("<end SimpleRange>n")
}
