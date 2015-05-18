// tar.go (c) David Rook 2010 - released under Simplified BSD 2-clause License

package main

import (
	"archive/tar"
	"flag"
	"fmt"
	"io"
	"os"
)

func list_headers(fname string) {
	f, err := os.Open(fname)
	if err != nil {
		fatal_err(err)
	}
	defer f.Close()

	tr := tar.NewReader(f)
	nfiles := 0
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			fatal_err(err)
		}
		if hdr == nil {
			break
		}
		fmt.Printf("%s contains: %s\n", fname, hdr.Name)
		nfiles++
	}
	fmt.Printf("Archive contains %d file headers\n", nfiles)
}

func main() {
	fmt.Printf("Testing tar.go\n")
	flag.Parse()
	if flag.NArg() == 0 { // do nothing
		fmt.Printf("? NO args - nothing to do...\n")
	} else {
		for i := 0; i < flag.NArg(); i++ {
			fmt.Printf("arg[%d] %s\n", i, flag.Arg(i))
			tarname := flag.Arg(i)
			list_headers(tarname)
		}
	}
}

func fatal_err(erx error) {
	fmt.Printf("%s \n", erx)
	os.Exit(1)
}
