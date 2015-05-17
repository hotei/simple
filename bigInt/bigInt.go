// bigInt.go

package main

import (
	"fmt"
	"math/big"
)

func gcd1(a, b *big.Int) *big.Int {
	for b.Sign() != 0 {
		a, b = b, new(big.Int).Mod(a, b)
	}
	return a
}

var newB *big.Int = new(big.Int)

func gcd2(a, b *big.Int) *big.Int {
	b = newB.Set(b)
	for b.Sign() != 0 {
		a, b = b, a.Mod(a, b)
	}
	return a
}

func rho(n *big.Int) *big.Int {
	// can simplify with a := big.NewInt(2)
	a := new(big.Int)
	b := new(big.Int)
	d := new(big.Int)
	one := new(big.Int)
	t := new(big.Int)

	a.SetString("2", 10)
	b.SetString("2", 10)
	d.SetString("1", 10)
	one.SetString("1", 10)
	// Python
	//    while d == 1:
	//      a = (a**2 + 1) % n
	//        b = (b**2 + 1) % n
	//        b = (b**2 + 1) % n
	//        d = gcd(a-b,n)
	//    return d

	for {
		if d.Cmp(one) != 0 {
			break
		}
		// a = (a*a + 1) % n
		a.Mul(a, a)
		a.Add(a, one)
		a.Mod(a, n)

		// b = (b*b + 1) % n
		b.Mul(b, b)
		b.Add(b, one)
		b.Mod(b, n)

		b.Mul(b, b)
		b.Add(b, one)
		b.Mod(b, n)

		t.Sub(a, b)
		if t.Sign() < 0 {
			t.Add(t, n)
		}
		d = gcd2(t, n)
		// d.GCD(nil,nil,t,n)
	}
	return d
}

func main() {
	b := new(big.Int)
	b.SetString("62823675885202669644104829577", 10)
	fmt.Printf("factors of %s are %v\n", b.String(), rho(b))
}
