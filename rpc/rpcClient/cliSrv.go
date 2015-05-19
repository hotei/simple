// cliSrv.go  (c) David Rook 2012 - released under Simplified BSD 2-clause License

/*
 * adapted from rpc example golang.org/pkg/net/rpc
 * usage :
 * go build rpcDemo package first
 *
 * go build
 * rm server ; ln rpcClient server
 * mv rpcClient client
 * ./server &
 * ./client
 *
 */

package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"time"
	//
	server "github.com/hotei/simple/rpc/rpcDemo"
)

var (
	// serverAddress = "10.1.2.115"	// odin
	serverAddress = "127.0.0.1" // myself for test
)

func FatalError(err error) {
	fmt.Printf("!Err-> %s\n", err)
	os.Exit(1)
}

func main() {
	fmt.Printf("command name = %s\n", os.Args[0])
	flag.Parse()

	if os.Args[0] == "./server" { // TODO pry off the ./ portion
		fmt.Printf("rpcServer starting at tcp:1234 will run 10.0 minutes only\n")
		arith := new(server.Arith)
		rpc.Register(arith)
		rpc.HandleHTTP()
		listener, e := net.Listen("tcp", ":1234")
		if e != nil {
			log.Fatal("listen error:", e)
		}
		go http.Serve(listener, nil)
		time.Sleep(10 * time.Minute)
		fmt.Printf("rpcServer ending service for tcp:1234\n")
	}
	if os.Args[0] == "./client" { // TODO pry off the ./ portion
		fmt.Printf("rpcClient starting\n")
		client, err := rpc.DialHTTP("tcp", serverAddress+":1234")
		if err != nil {
			log.Fatal("dialing:", err)
		}
		// Synchronous call
		var args = server.Args{A: 7, B: 8}
		var reply int
		err = client.Call("Arith.Multiply", args, &reply)
		if err != nil {
			log.Fatal("arith error:", err)
		}
		fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)

		var reply2 server.Quotient
		args = server.Args{A: 52, B: 10}
		err = client.Call("Arith.Divide", args, &reply2)
		if err != nil {
			log.Fatal("arith error:", err)
		}
		fmt.Printf("Arith: %d div(%d) = %d, mod(%d)= %d\n", args.A, args.B, reply2.Quo, args.B, reply2.Rem)

		// Asynchronous call
		quotient := new(server.Quotient)
		divCall := client.Go("Arith.Divide", args, quotient, nil)
		_ = <-divCall.Done
		if divCall.Error != nil {
			log.Fatal("arith error:", err)
		}
		fmt.Printf("divCall(%v)\n", divCall)
		fmt.Printf("divCall.Error(%v)\n", divCall.Error)
		fmt.Printf("quotient(%v)  Quo(%d) Rem(%d)\n", quotient, quotient.Quo, quotient.Rem)
		fmt.Printf("rpcClient ending\n")
	}
}

/*  copy of the pkg in case files get separated

package rpcDemo

import (
	"errors"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}
*/
