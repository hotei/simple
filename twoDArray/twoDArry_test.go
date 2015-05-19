// Simple2DArray_test.go

package main

import (
	"fmt"
	"log"
	"testing"
)

func Test_01(t *testing.T) {
	debug := false
	k := 0
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			val := twoDarray[i][j]
			if debug {
				fmt.Printf("twoDarray[%d][%d] = %d\n", i, j, val)
			}
			if val != k {
				t.Errorf("fail Test_01()")
			}
			k++
		}
	}
	if false {
		fmt.Printf("Test\n")
		log.Panic("Panic on purpose to inspect code")
		t.Errorf("forced fail in Test_Env()")
	}
}
