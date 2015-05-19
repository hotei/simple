// rpcDemo.go (c) David Rook 2012 - released under Simplified BSD 2-clause License

/*
 * copy of golang.org/pkg/net/rpc example
 *
 * usage :
 * go build rpcDemo package first (this file)
 *
 * switch to rpcClient directory
 * note: same code for client & server
 * go build
 * rm server ; ln rpcClient server
 * mv rpcClient client
 * ./server &
 * ./client
 *
 */

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
