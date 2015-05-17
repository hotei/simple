// call_test.go (c) 2010 David Rook

/*
 *
 */

package main

import (
	"fmt"
	"os"
	"testing"
)

func Test_Call(t *testing.T) {
	debug = false
	if debug {
		fmt.Printf("Running Test_Call()\n")
	}
	summit(10) // no return value here
	x := summitAndReturn(10)
	if x != 45 {
		t.Errorf("summitAndReturn(10) expected 45, got(%d)", x)
	}
	if debug {
		fmt.Printf("summitAndReturn(10) gives return of %d\n", x)
	}
	who_Am_I_now()
}

// working ok
func Test_who_Am_I(t *testing.T) {
	debug = false
	host, errCode := os.Hostname()
	if errCode != nil {
		t.Errorf("os.Hostname() Error %v", errCode)
	}
	if debug {
		fmt.Printf("2 Host = %s\n", host)
	}
}
