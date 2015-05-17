// closure.go

package main

import (
	"fmt"
	"os"
	"time"
)

func makecounter(c int) func() int { return func() int { c++; return c } }

func test1() {
	f := makecounter(0)
	g := makecounter(0)
	fmt.Printf("f(%d) g(%d)\n", f(), g())
	fmt.Printf("f(%d)\n", f())
	fmt.Printf("f(%d) g(%d)\n", f(), g())
}

// these would normally be C static vars
var spinCt int8
var spinchars string = "|/-\\|/-\\ "

func spinner() {
	fmt.Fprintf(os.Stderr, "%s\r", spinchars[spinCt:spinCt+1])
	spinCt++
	spinCt &= 0x7 // mod 8 which is length of spinchars by design
}

func spin3(spinCt int) {
	spinCt &= 0x7 // mod 8 which is length of spinchars by design
	fmt.Fprintf(os.Stderr, "%s\r", spinchars[spinCt:spinCt+1])
}

func test2() {
	sec := 3
	fmt.Printf("test2 - show spinner for %d seconds\n", sec)
	for i := 0; i < 1000; i++ {
		spinner()
		time.Sleep(time.Duration(sec) * time.Millisecond)
	}
	fmt.Printf("<done>\n")
}

func makespinner(c int) func() int { return func() int { c++; return c } }

func test3() {
	f := makespinner(0)
	fmt.Printf("%d\n", f())
	fmt.Printf("%d\n", f())

	fmt.Printf("test 3 - show spinner for 10 seconds\n")
	for i := 0; i < 1000; i++ {
		spin3(f())
		time.Sleep(10 * time.Millisecond)
	}
	fmt.Printf("<done>\n")

}

func main() {
	test2()
	test3()
}
