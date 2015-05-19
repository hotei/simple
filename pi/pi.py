#! /usr/bin/env python

import sys

def main():
    j=0
    enough = 40
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
