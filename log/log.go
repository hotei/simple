// log.go (c) David Rook 2010 - released under Simplified BSD 2-clause License

package main

import (
	"fmt"
	"log"
	"os"
)

func callme2() {
	log.SetPrefix("callme2: ")
	log.Printf("second dated log entry with line number")
}

func main() {
	status := 0

	tmp := os.Getenv("SHELL")
	if tmp != "" {
		fmt.Printf("Shell in use is %s\n", tmp)
	}

	log.SetPrefix("main: ")
	log.SetFlags(log.Ldate | log.Llongfile)
	log.SetOutput(os.Stdout)
	log.Printf("first dated log entry with file & line number")
	callme2()
	// note next call to log.Printf( will have callme2 as prefix unless you change it back
	log.SetPrefix("main: ")
	log.Printf("third dated log entry with file & line number")

	fmt.Printf("SimpleLog exiting\n")
	os.Exit(status)
}
