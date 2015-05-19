// scrypter.go (c) David Rook 2010 - released under Simplified BSD 2-clause License

package main

import (
	"fmt"
	"math/rand"
	//
	"code.google.com/p/go.crypto/scrypt"
	//
	"github.com/hotei/mdr"
)

// seed math.rand() generator and return the value used as seed
func RandomSeed() []byte {
	//	seed := mdr.UdevRandomBlock(8)
	seed := []byte{0, 0, 0, 0, 0, 0, 0, 1}
	seedInt64 := mdr.Int64FromLSBytes(seed)
	rand.Seed(seedInt64)
	return seed
}

// runtime is roughly proportional to r * p which must be less than 1 << 30
// this is intended to be SLOW, to foil brute force password attacks.
func main() {
	password := []byte("this is my key")
	salt := mdr.UdevRandomBlock(8)
	N := 16384
	r := 8
	p := 1
	keyLen := 32
	newKey, err := scrypt.Key(password, salt, N, r, p, keyLen)
	if err != nil {
		mdr.FatalError(err)
	}
	fmt.Printf("newKey : %x\n", newKey)
	fmt.Printf("r * p (%d) < (%d) \n", r*p, 1<<30)
	fmt.Printf("done\n")
}
