// memAvail.go (c) David Rook 2012 - released under Simplified BSD 2-clause License

package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func MemAvail() int64 {
	fname := "/proc/meminfo"
	FileBytes, err := ioutil.ReadFile(fname)
	if err != nil {
		fmt.Printf("!Err-> can't read %s\n", fname)
		return -6
	}
	bufr := bytes.NewBuffer(FileBytes)
	for {
		line, err := bufr.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Sprintf("!Err-> can't read %s\n", fname)
		}
		ndx := strings.Index(line, "MemFree:")
		if ndx >= 0 {
			line = strings.TrimSpace(line[9:])
			fmt.Printf("%q\n", line)
			line = line[:len(line)-3]
			fmt.Printf("%q\n", line)
			mem, err := strconv.ParseInt(line, 10, 64)
			if err == nil {
				return mem
			}
			// some problem if parse failed
			fmt.Printf("line: %s\n", line)
			n, err := fmt.Sscan(line, "%d", &mem)
			if err != nil {
				fmt.Printf("!Err-> can't scan %s\n", line)
				return -2
			}
			if n != 1 {
				fmt.Printf("!Err-> can't scan all %s\n", line)
				return -3
			}
			return -4
		}
	}
	fmt.Printf("didn't find MemFree in /proc/meminfo\n")
	return -5
}

func FatalError(err error) {
	fmt.Printf("!Err-> %s\n", err)
	os.Exit(1)
}

func main() {
	fmt.Printf("hello\n")
	fmt.Printf("MemAvail(%d)\n", MemAvail())
}
