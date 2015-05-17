// treewalk.go (c) 2010 David Rook - released with Simplified BSD 2-clause Lincense

// list the dir tree for each comand arg

package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
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
