// filter.go

// This is supposed to gather stdin into a []string slice
// then copy the slice to a file named output.lst
// use ^D to terminate input from keyboard

//		16 GB RAM, Ubuntu 11.04/AMD64
//			go Weekly 2012-03-13

// working ok with go 1.4.2

package main

import (
	"fmt"
	"github.com/hotei/mdr"
	"os"
)

func main() {
	argList := mdr.GetAllArgs()
	fp, err := os.Create("output.lst")
	if err != nil {
		fmt.Printf("Cant create output.lst\n")
		os.Exit(-1)
	}
	defer fp.Close()
	for _, s := range argList {
		s = fmt.Sprintf("%s\n", s)
		nw, err := fp.WriteString(s)
		if (err != nil) || (nw <= 0) { // never 0 on success
			fmt.Printf("Cant write output.lst\n")
			break
		}
	}
	fmt.Printf("Args gathered = %v\n", argList)
}
