// sleeper.go (c) David Rook 2010 - released under Simplified BSD 2-clause License

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func one(ch chan int64) {
	fmt.Printf("test_one\n")
	x := rand.Int63n(10)
	time.Sleep(time.Duration(x) * time.Second)
	ch <- x
}

func main() {
	fmt.Printf("Main started...\n")
	now := time.Now()
	nowt := now.Unix()
	rand.Seed(nowt)

	ch := make(chan int64, 1)
	timeout := make(chan bool, 1)

	// must start the timeout function first
	go func() {
		time.Sleep(3 * time.Second) // 3 seconds
		timeout <- true
		// fmt.Printf("timeout inside block\n")
	}()

	// then call the test function
	go one(ch)

	// now see which one finishes first
	var x int64
	select {
	case x = <-ch: // read occurred here
		fmt.Printf("read from test subr number %d\n", x)
		//break

	case <-timeout: // timed out
		fmt.Printf("Timeout reached\n")
		//break
	}
	fmt.Printf("Main finished\n")
}

/* below is not mine, probably from golang.org documents

func Query(conns []Conn, query string) Result {
    ch := make(chan Result, 1)
    for _, conn := range conns {
        go func(c Conn) {
            _ = ch <- c.DoQuery(query)
        }(conn)
    }
    return <- ch
}

In this example, the closure does a non-blocking send, which it achieves by using the send
operation in an expression; the success/failure of the send is a boolean value that is the
result of the expression. Here the value is just tossed away by assigning it to the blank
identifier '_', but that's sufficient to prevent the send from blocking. Making the send
non-blocking guarantees that none of the goroutines launched in the loop will hang around.

However, if the result arrives before the main function has made it to the receive,
the send could fail since no one is ready.

This problem is a textbook of example of what is known as a race condition, but the fix is trivial.
We just make sure to buffer the channel ch (by adding the buffer length as the second argument
to make), guaranteeing that the first send has a place to put the value. This ensures the send
will always succeed, and the first value to arrive will be retrieved regardless of the order
of execution.

*/
