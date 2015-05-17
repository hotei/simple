// get OS to tell you a free port you can use for server

package main

import (
	"net"
)

func main() {

	l, _ := net.Listen("tcp", ":0")
	defer l.Close()
	println(l.Addr().String())

}
