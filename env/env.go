// env.go (c) 2010 David Rook - released with Simplified BSD 2-clause Lincense

package main

import (
	"fmt"
	"os"
)

func main() {

	tmp := os.Getenv("SHELL")
	if tmp != "" {
		fmt.Printf("Shell in use is: %s\n", tmp)
	} else {
		fmt.Printf("Unexpected error\n")
		os.Exit(-1)
	}

	tmp = os.Getenv("DISPLAY")
	if tmp != "" {
		fmt.Printf("X-11 Display in use is: %s\n", tmp)
	} else {
		fmt.Printf("Unexpected error\n")
		os.Exit(-1)
	}

	tmp = os.Getenv("NOTADISPLAY")
	if tmp == "" {
		fmt.Printf("NOTADISPLAY is not an env variable\n")
	} else {
		fmt.Printf("spurious environment returned\n")
		os.Exit(-1)
	}

	//for ndx, val := range os.Envs {		// no longer available in GO-1
	//fmt.Printf("ENV %d = %v\n", ndx, val)
	//}
	x := os.Environ()
	i := 0
	for _, ea := range x {
		fmt.Printf("Env[%d] = %s\n", i, ea)
		i++
	}
}
