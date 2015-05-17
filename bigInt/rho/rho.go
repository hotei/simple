// rho.go
package rho

import (
	"math/big"
)

func FromString(s string) *big.Int {
	val, _ := new(big.Int).SetString(s, 0)
	return val
}

/*
#Python version
def rho(n):
    a = 2
    b = 2
    d = 1
    while d == 1:
        a = (a**2 + 1) % n
        b = (b**2 + 1) % n
        b = (b**2 + 1) % n
        d = gcd(a-b,n)
    return d
*/

func Rho(n *big.Int) *big.Int {
	a := big.NewInt(2)
	b := big.NewInt(2)
	d := big.NewInt(1)
	z := big.NewInt(0)
	one := big.NewInt(1)

	for d.Cmp(one) == 0 {
		a.Mul(a, a)
		a.Add(a, one)
		a.Mod(a, n)

		b.Mul(b, b)
		b.Add(b, one)
		b.Mod(b, n)

		b.Mul(b, b)
		b.Add(b, one)
		b.Mod(b, n)

		z.Sub(a, b)
		if z.Sign() < 0 {
			z.Add(z, n)
		}

		d.GCD(nil, nil, z, n)
	}

	return d
}
