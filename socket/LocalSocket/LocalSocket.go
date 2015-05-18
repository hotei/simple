// LocalSocket.go (c) David Rook 2010 - released under Simplified BSD 2-clause License

/*
 * this corresponds roughly to client1/server1 in vwar/test_socket (working)

 TODO net offers a timeout function for listener
 TODO or a non-blocking channel,but that's polling and not optimal
*/

package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

// writes one character to the server, then gets back a result which we can check for validity
func client_A(d byte, done chan int) {
	ch_out := []byte{d}
	ch_in := []byte{d}

	c, err := net.Dial("unix", "server_socket")
	_, err = c.Write(ch_out)
	if err != nil {
		fmt.Printf("client write to socket failed with err: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("client wrote byte value = %d to socket\n", ch_out)

	_, err = c.Read(ch_in)
	if err != nil {
		fmt.Printf("client read from socket failed with err: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("client read byte value = %d from socket\n", ch_in)
	if ch_in[0] != (ch_out[0] * 10) {
		fmt.Printf("didn't get expected result\n")
	}
	_ = c.Close()
	done <- 1
}

// how to tell server to close listener without blocking or polling?  TODO
func server_A(ready chan int, quit chan int) {
	ch := []byte{0}
	var n int
	var listener net.UnixAddr
	listener.Name = "server_socket"

	_, err := os.Open(listener.Name)
	if err == nil {
		fmt.Printf("socket already in use\n")
		os.Exit(-1)
	}
	l, err := net.ListenUnix("unix", &listener)
	if err != nil {
		fmt.Printf("ListenUnix()  failed with err: %v\n", err)
		os.Exit(1)
	}
	defer l.Close()
	t := time.Now()
	t = t.Add(10 * time.Second)
	err = l.SetDeadline(t)
	fmt.Printf("server() listener started ok, ready to accept connections\n")
	ready <- 1 // send ready signal back to main()
	for {
		//select {  didn't solve problem
		//case _ = <- quit: break
		//default:
		//}
		c, err := l.AcceptUnix()
		if err != nil {
			fmt.Printf("server() accept socket failed with err: %v\n", err)
			break
		}
		fmt.Printf("server() accepted connection from client\n")
		n, err = c.Read(ch)
		if err != nil || n != 1 {
			fmt.Printf("server() read socket failed with err: %v\n", err)
			os.Exit(1)
		}
		ch[0] *= 10 // modify input data and send it back
		n, err = c.Write(ch)
		if err != nil || n != 1 {
			fmt.Printf("server() write to socket failed with err: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("server() wrote byte value = %d to socket\n", ch)
		c.Close()
	}
	fmt.Printf("Cleaning up after listener\n")
	err = os.Remove(listener.Name)
	if err != nil {
		log.Panicf("cant remove socket")
	}
	quit <- 1
}

func main() {
	flag.Parse()
	if false {
		log.Panicf("don't")
	}
	ready := make(chan int, 1)
	done := make(chan int, 1)
	quit := make(chan int, 1)
	fmt.Printf("main() starting 1 LocalSocket server\n")
	go server_A(ready, quit)
	_ = <-ready // wait for server to initialize
	fmt.Printf("main() server says it's ready, time to start clients...\n")
	fmt.Printf("main() starting 1 LocalSocket client(s)\n")
	workers := 10
	for i := 0; i < workers; i++ {
		go client_A(byte(i), done)
	}
	waitfor := workers
	for waitfor > 0 {
		_ = <-done
		waitfor--
	}
	_ = <-quit
	//time.Sleep(11 * time.Second)
	fmt.Printf("main() exit\n")
}
