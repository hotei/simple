package main

import (
	"fmt"
	"math/big"
	"testing"
)

func Test_N5(t *testing.T) {
	fmt.Printf("Test_N5...\n")
	a := new(big.Int)
	a.SetString("766150476015982127183457373", 10)
	r := rho(a)
	fmt.Printf("gcd factor of %s is %v\n", a.String(), r)
	if false {
		rr := rho(r)
		fmt.Printf("gcd factor of %s is %v\n", rr.String(), rr)
		if r.Cmp(rr) == 0 {
			fmt.Printf("%s is prime\n", rr.String())
		}
	}
	// should be 1178524040059
}

func Test_N6(t *testing.T) {
	fmt.Printf("Test_N6...\n")
	a := new(big.Int)
	a.SetString("62823675885202669644104829577", 10)
	fmt.Printf("gcd factor of %s is %v\n", a.String(), rho(a))

	// should be 663410067979
}
