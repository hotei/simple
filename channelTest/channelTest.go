// channelTest.go  from ?golang.org perhaps?

/*
Simple example of taking turns reading and writing without locks
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	wg sync.WaitGroup
)

func reader(chin, chout chan int) {
	wg.Add(1)
	defer wg.Done()
	leftToDo := 3
	for {
		x := <-chin
		if x < 0 {
			return
		}

		fmt.Printf("reading - left to do = %d\n", leftToDo)
		leftToDo--
		time.Sleep(1 * time.Second)
		if leftToDo <= 0 {
			// signal writer that no more data is expected
			chout <- 0
			return
		}
		chout <- x
	}
}

func writer(chin, chout chan int) {
	wg.Add(1)
	defer wg.Done()
	for {
		x := <-chin
		if x < 0 {
			return
		}
		fmt.Printf("Writing - left to do = %d\n", x)
		time.Sleep(2 * time.Second)
		if x == 0 {
			return // got last packet from reader
		}
		chout <- x
	}
}

func doit() {
	chin := make(chan int)
	chout := make(chan int)
	go reader(chin, chout)
	go writer(chout, chin)
	chin <- 1
	wg.Wait()
}

func main() {
	fmt.Printf("go\n")
	doit()
}
