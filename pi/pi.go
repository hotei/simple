// pi.go (c) David Rook 2010 - released under Simplified BSD 2-clause License

/*
 * Sortof works.  Gives up at about 9 digits
 * See SimplePi.py for what it SHOULD look like
 * First 40 digits from python 31415926535897932384626433832795028841971
 */

package main

import (
	"fmt"
)

func makePi(n int) {
	fmt.Printf("%d digits of pi have been requested\n", n)
	j := 0
	enough := n
	k := int64(2)
	a := int64(4)
	b := int64(1)
	a1 := int64(12)
	b1 := int64(4)
	var p, q, d, d1 int64

	for {
		p, q, k = k*k, 2*k+1, k+1
		a, b, a1, b1 = a1, b1, p*a+q*a1, p*b+q*b1
		d, d1 = a/b, a1/b1
		for d == d1 {
			j++
			output(uint8(d))
			a, a1 = 10*(a%b), 10*(a1%b1)
			d, d1 = a/b, a1/b1
		}
		if j > enough {
			break
		}
	}
	fmt.Println()
}

func output(d byte) {
	fmt.Printf("%c", byte('0')+d)
}

func main() {
	fmt.Printf("SimplePi.go\n")
	for i := 3; i < 15; i++ {
		makePi(i)
	}
}

/*
#! /usr/bin/env python
import sys

def main():
   j=0
    enough = 40000
    k, a, b, a1, b1 = 2L, 4L, 1L, 12L, 4L
    while 1:
        p, q, k = k*k, 2L*k+1L, k+1L
        a, b, a1, b1 = a1, b1, p*a+q*a1, p*b+q*b1
        d, d1 = a/b, a1/b1
        while d == d1:
	    j=j+1
            output(d)
            a, a1 = 10L*(a%b), 10L*(a1%b1)
            d, d1 = a/b, a1/b1
	if j > enough: break
    print

def output(d):
    sys.stdout.write(`int(d)`)
    sys.stdout.flush()

main()

*/
