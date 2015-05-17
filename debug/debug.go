package main

import (
	"flag"
	"fmt"
)

var (
	verboseFlag bool
)

type Debug bool

func init() {
	flag.BoolVar(&verboseFlag, "v", false, "be verbose")
}

func (d Debug) Printf(s string, a ...interface{}) {
	if d {
		fmt.Printf(s, a...)
	}
}

func main() {
	var dbg Debug
	flag.Parse()
	dbg = Debug(verboseFlag)
	dbg.Printf("Debugging set by verboseFlag\n")
	dbg = true
	dbg.Printf("Debugging on\n")
	dbg = false
	dbg.Printf("Debugging off\n")

}
