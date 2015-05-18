// pipe.go (c) David Rook 2010 - released under Simplified BSD 2-clause License

package main

import (
	"fmt"
	"os"
	"time"
)

func reading(reader *os.File) {
	const NBUF = 512
	buf := make([]byte, NBUF)

	for {
		r, _ := reader.Read(buf)
		if r > 0 {
			fmt.Printf("read this in: %s\n", buf[0:r])
		}
	}
}

func writing(writer *os.File) {
	teststring := "test"
	fmt.Printf("wrote this to pipe: %s\n", teststring)
	writer.WriteString(teststring)
	time.Sleep(1 * 1e9)
	teststring = "stringtest"
	fmt.Printf("wrote this to pipe: %s\n", teststring)
	writer.WriteString(teststring)
}

func main() {

	fmt.Printf("<start pipe>\n")

	rdr, wrt, err := os.Pipe()

	if err != nil {
		fmt.Printf("error of some sort %v\n", err)
		os.Exit(1)
	}
	go writing(wrt)
	go reading(rdr)

	time.Sleep(3 * 1E9)

	fmt.Printf("<end pipe>\n")
}
