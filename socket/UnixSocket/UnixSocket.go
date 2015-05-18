// UnixSocket.go (c) David Rook 2010 - released under Simplified BSD 2-clause License

package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	//    "time"
)

// writes one character to the server, then gets back a result which we can check for validity
func client_A(d byte, done chan int) {
	ch_out := []byte{d}
	ch_in := []byte{d}

	c, err := net.Dial("tcp", "127.0.0.1:8888")
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

func server_A(ready chan int) {
	ch := []byte{0}
	var n int
	var listener net.TCPAddr
	listener.IP = net.ParseIP("127.0.0.1")
	listener.Port = 8888

	l, err := net.ListenTCP("tcp4", &listener)
	if err != nil {
		fmt.Printf("ListenTCP()  failed with err: %v\n", err)
		os.Exit(1)
	}
	defer l.Close()
	fmt.Printf("server() listener started ok, ready to accept TCP connections\n")
	ready <- 1 // send ready signal back to main()
	for {
		c, err := l.AcceptTCP()
		if err != nil {
			fmt.Printf("server() accept TCP failed with err: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("server() accepted TCP connection from client\n")
		n, err = c.Read(ch)
		if err != nil || n != 1 {
			fmt.Printf("server() read TCP failed with err: %v\n", err)
			os.Exit(1)
		}
		ch[0] *= 10 // modify input data and send it back
		n, err = c.Write(ch)
		if err != nil || n != 1 {
			fmt.Printf("server() write to TCP failed with err: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("server() wrote byte value = %d to TCP\n", ch)
		c.Close()
	}
}

func main() {
	flag.Parse()
	if false {
		log.Panicf("don't")
	}
	ready := make(chan int, 1)
	quit := make(chan int, 1)

	fmt.Printf("main() starting 1 TCP server\n")
	go server_A(ready)
	_ = <-ready // wait for server to initialize
	fmt.Printf("main() TCP server says it's ready, time to start clients...\n")
	fmt.Printf("main() starting 1 TCP client(s)\n")
	workers := 10
	for i := 0; i < workers; i++ {
		go client_A(byte(i), quit)
	}
	waitfor := workers
	for waitfor > 0 {
		_ = <-quit
		waitfor--
	}
	fmt.Printf("main() exit\n")
}
