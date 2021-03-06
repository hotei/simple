//scrypt_test.go (c) David Rook 2010 - released under Simplified BSD 2-clause License

package main

import (
	// standard pkgs only below
	"bytes"
	"fmt"
	"testing"
	//
	"code.google.com/p/go.crypto/scrypt"
	//
	"github.com/hotei/mdr"
)

func Benchmark_001(b *testing.B) {
	password := []byte("this is my key")
	salt := mdr.UdevRandomBlock(8)
	N := 16384
	r := 8
	p := 1
	keyLen := 32
	for i := 0; i < b.N; i++ {
		_, err := scrypt.Key(password, salt, N, r, p, keyLen)
		if err != nil {
			mdr.FatalError(err)
		}
	}
}

func Test_001(t *testing.T) {
	fmt.Printf("Test_001 running...\n")
	password := []byte("")
	salt := []byte("")
	N := 16
	r := 1
	p := 1
	keyLen := 64
	newKey, err := scrypt.Key(password, salt, N, r, p, keyLen)
	if err != nil {
		mdr.FatalError(err)
	}
	fmt.Printf("newKey = %x\n", newKey)
	fmt.Printf("r * p (%d) < (%d) \n", r*p, 1<<30)
	var referenceKey = []byte{
		0x77, 0xd6, 0x57, 0x62, 0x38, 0x65, 0x7b, 0x20, 0x3b, 0x19, 0xca, 0x42, 0xc1, 0x8a, 0x04, 0x97,
		0xf1, 0x6b, 0x48, 0x44, 0xe3, 0x07, 0x4a, 0xe8, 0xdf, 0xdf, 0xfa, 0x3f, 0xed, 0xe2, 0x14, 0x42,
		0xfc, 0xd0, 0x06, 0x9d, 0xed, 0x09, 0x48, 0xf8, 0x32, 0x6a, 0x75, 0x3a, 0x0f, 0xc8, 0x1f, 0x17,
		0xe8, 0xd3, 0xe0, 0xfb, 0x2e, 0x0d, 0x36, 0x28, 0xcf, 0x35, 0xe2, 0x0c, 0x38, 0xd1, 0x89, 0x06}
	if bytes.Compare(referenceKey, newKey) != 0 {
		t.Errorf("Test_001 failed")
	}
}

func Test_004(t *testing.T) {
	fmt.Printf("Test_004 running...\n")
	password := []byte("pleaseletmein")
	salt := []byte("SodiumChloride")
	N := 1048576
	r := 8
	p := 1
	keyLen := 64
	newKey, err := scrypt.Key(password, salt, N, r, p, keyLen)
	if err != nil {
		mdr.FatalError(err)
	}
	fmt.Printf("newKey = %x\n", newKey)
	fmt.Printf("r * p (%d) < (%d) \n", r*p, 1<<30)
	var referenceKey = []byte{
		0x21, 0x01, 0xcb, 0x9b, 0x6a, 0x51, 0x1a, 0xae, 0xad, 0xdb, 0xbe, 0x09, 0xcf, 0x70, 0xf8, 0x81,
		0xec, 0x56, 0x8d, 0x57, 0x4a, 0x2f, 0xfd, 0x4d, 0xab, 0xe5, 0xee, 0x98, 0x20, 0xad, 0xaa, 0x47,
		0x8e, 0x56, 0xfd, 0x8f, 0x4b, 0xa5, 0xd0, 0x9f, 0xfa, 0x1c, 0x6d, 0x92, 0x7c, 0x40, 0xf4, 0xc3,
		0x37, 0x30, 0x40, 0x49, 0xe8, 0xa9, 0x52, 0xfb, 0xcb, 0xf4, 0x5c, 0x6f, 0xa7, 0x7a, 0x41, 0xa4}

	if bytes.Compare(referenceKey, newKey) != 0 {
		t.Errorf("Test_004 failed")
	}
}

func Benchmark_005a(b *testing.B) {
	//	fmt.Printf("r * p (%d) < (%d) \n", r*p, 1<<30)
	password := []byte("")
	salt := []byte("")
	N := 16
	r := 1
	p := 1
	keyLen := 64
	var referenceKey = []byte{
		0x77, 0xd6, 0x57, 0x62, 0x38, 0x65, 0x7b, 0x20, 0x3b, 0x19, 0xca, 0x42, 0xc1, 0x8a, 0x04, 0x97,
		0xf1, 0x6b, 0x48, 0x44, 0xe3, 0x07, 0x4a, 0xe8, 0xdf, 0xdf, 0xfa, 0x3f, 0xed, 0xe2, 0x14, 0x42,
		0xfc, 0xd0, 0x06, 0x9d, 0xed, 0x09, 0x48, 0xf8, 0x32, 0x6a, 0x75, 0x3a, 0x0f, 0xc8, 0x1f, 0x17,
		0xe8, 0xd3, 0xe0, 0xfb, 0x2e, 0x0d, 0x36, 0x28, 0xcf, 0x35, 0xe2, 0x0c, 0x38, 0xd1, 0x89, 0x06}

	for i := 0; i < b.N; i++ {
		newKey, err := scrypt.Key(password, salt, N, r, p, keyLen)
		if err != nil {
			mdr.FatalError(err)
		}
		// fmt.Printf("newKey = %x\n", newKey)
		if bytes.Compare(referenceKey, newKey) != 0 {
			b.Errorf("Test_001 failed")
		}
	}
}

func Benchmark_005b(b *testing.B) {
	//	fmt.Printf("r * p (%d) < (%d) \n", r*p, 1<<30)
	password := []byte("")
	salt := []byte("")
	N := 16
	r := 10
	p := 10
	keyLen := 64

	for i := 0; i < b.N; i++ {
		_, err := scrypt.Key(password, salt, N, r, p, keyLen)
		if err != nil {
			mdr.FatalError(err)
		}
		// fmt.Printf("newKey = %x\n", newKey)
	}
}
