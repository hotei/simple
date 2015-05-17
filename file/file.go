// file.go (c) David Rook 2010 - released under Simplified BSD 2-clause License

package main

import (
	"fmt"
	"os"
)

func simpleReader(fname string) {
	file, err := os.Open(fname)
	if file == nil {
		fmt.Printf("can't open file %s as readonly; err=%v\n", fname, err)
		os.Exit(1)
	}
	fmt.Printf("Opened %s without error\n", fname)
	fmt.Printf("Sending %s to stdout\n", fname)
	const NBUF = 512

	// var buf [NBUF]byte
	buf := make([]byte, NBUF)
	for {
		nr, er := file.Read(buf)
		switch {
		case nr < 0:
			fmt.Fprintf(os.Stderr, "cat: error reading from %s: %v\n", fname, er)
			os.Exit(1)
		case nr == 0: // EOF
			return
		case nr > 0:
			nw, ew := os.Stdout.Write(buf[0:nr])
			if nw != nr {
				fmt.Fprintf(os.Stderr, "cat: error writing from %s: %v\n", fname, ew)
			}
		}
	}
}

func simpleWriter(fname string) {
	file, err := os.Create(fname)
	if file == nil {
		fmt.Printf("can't open file for read/write; err=%v\n", err)
		os.Exit(1)
	}
	file.WriteString("test 1\n")
	file.WriteString("test 2\n")
	err = file.Close()
}

func main() {
	fmt.Printf("<SimpleFile.go starting>\n")

	simpleReader("./file.go")
	simpleWriter("./testfile.txt")
	fmt.Printf("Writing out testfile.txt\n")
	fmt.Printf("Reading in our new testfile.txt\n")
	simpleReader("./testfile.txt")
	fmt.Printf("<file.go finished>\n")
}
