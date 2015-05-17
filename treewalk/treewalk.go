// treewalk.go (c) 2010 David Rook - released with Simplified BSD 2-clause Lincense

// list the dir tree for each comand arg

package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func ls(wpath string, info os.FileInfo, err error) error {
	if info == nil {
		fmt.Printf("no stat info available for %s\n", wpath)
		return nil
	}
	if info.IsDir() {
		// fmt.Printf("dir %s\n",path)
	} else {
		fmt.Printf("file %q size(%d)\n", wpath, info.Size())
	}
	return nil
}

func main() {
	flag.Parse()
	fmt.Printf("NArg = %d\n", flag.NArg())
	if flag.NArg() < 1 {
		fmt.Printf("quitting - nothing to do\n")
	} else {
		for i := 0; i < flag.NArg(); i++ {
			filepath.Walk(flag.Arg(i), ls)
		}
	}
}
