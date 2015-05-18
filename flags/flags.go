// flags.go (c) David Rook 2010 - released under Simplified BSD 2-clause License

// get arguments from command line

// ./flags a b c
// also try with an invalid flag - ./flags -v for instance
// if there are no valid flags you get a useless usage message
//      "Usage of ./flags:"

package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fmt.Printf("command name = %s\n", os.Args[0])
	flag.Parse()
	fmt.Printf("Flag got %d args on cmd line after command name\n", flag.NArg())
	if flag.NArg() == 0 { // do nothing
	} else {
		for i := 0; i < flag.NArg(); i++ {
			fmt.Printf("%d %s\n", i, flag.Arg(i))
		}
	}
}
