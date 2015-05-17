// call.go (c) 2010 David Rook

package main

import (
	"fmt"
	"log"
	"os"
)

var debug = true

// working ok
func summit(n int) {
	sum := 0
	for i := 0; i < n; i++ {
		sum += i
	}
	if debug {
		fmt.Printf("summit(%d) = %d\n", n, sum)
	}
}

// working ok
func summitAndReturn(n int) int {
	sum := 0
	for i := 0; i < n; i++ {
		sum += i
	}
	return sum
}

// working ok
func who_Am_I_now() {
	host, _ := os.Hostname() // note risky use of empty error return
	if debug {
		fmt.Printf("1 Host = %s\n", host)
	}
}

func main() {
	summit(10) // no return value here
	x := summitAndReturn(10)
	host, errCode := os.Hostname()
	if errCode != nil {
		log.Panic("don't know my own hostname")
	}
	if debug {
		fmt.Printf("%s %d \n", host, x)
	}
}
