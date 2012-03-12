// SimpleTreeWalk.go

// list the dir tree for each comand arg
/* 
 * (c) 2010 David Rook
 */

package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

var (
	license = "treewalk.go (c) 2012 David Rook released under Simplified BSD License"
)

func ls(path string, info os.FileInfo, err error) error {
	if info.IsDir() {
		// fmt.Printf("dir %s\n",path)
	} else {
		fmt.Printf("%s size(%d)\n", path, info.Size())
	}
	return nil
}

func main() {
	flag.Parse()
	if flag.NArg() <= 0 { // do nothing
	} else {
		for i := 0; i < flag.NArg(); i++ {
			filepath.Walk(flag.Arg(i), ls)
		}
	}
}
