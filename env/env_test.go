// env_test.go (c) 2010 David Rook - released with Simplified BSD 2-clause Lincense

package main

import (
	"fmt"
	"os"
	"testing"
)

func Test_Env(t *testing.T) {
	debug := true
	tmp := os.Getenv("SHELL")
	if tmp != "" {
		if debug {
			fmt.Printf("Shell in use is: %s\n", tmp)
		}
	} else {
		t.Errorf("SHELL env var not found")
	}

	tmp = os.Getenv("DISPLAY")
	if tmp != "" {
		if debug {
			fmt.Printf("X-11 Display in use is: %s\n", tmp)
		}
	} else {
		t.Errorf("DISPLAY env var not found")
	}

	tmp = os.Getenv("NOTADISPLAY")
	if tmp == "" {
		if debug {
			fmt.Printf("NOTADISPLAY is not an env variable\n")
		}
	} else {
		t.Errorf("spurious environment returned\n")

	}

	//for ndx, val := range os.Envs {		// no longer available in GO-1
	//fmt.Printf("ENV %d = %v\n", ndx, val)
	//}
	x := os.Environ()
	if len(x) <= 0 {
		t.Errorf("Empty environment")
	} else {
		if debug {
			i := 0
			for _, ea := range x {
				fmt.Printf("Env[%d] = %s\n", i, ea)
				i++
			}
		}
	}
	if false {
		t.Errorf("forced fail in Test_Env()")
	}
}
