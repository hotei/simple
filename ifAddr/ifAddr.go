// ifAddr.go (c) David Rook 2010 - released under Simplified BSD 2-clause License

/*
mdr@loki-535:~/Desktop/MYGO/src/Simple/ifAddr$ ./ifAddr
[127.0.0.1/8 10.1.2.112/24 ::1/128 fe80::16da:e9ff:feda:a313/64]
[{1 16436 lo  up|loopback} {2 1500 eth3 14:da:e9:da:a3:13 up|broadcast|multicast}]
*/

package main

import (
	"fmt"
	"net"
)

func main() {

	ifas, err := net.InterfaceAddrs()
	if err != nil {
		return
	}
	fmt.Printf("%v\n", ifas)

	ifs, err := net.Interfaces()
	if err != nil {
		return
	}
	fmt.Printf("%v\n", ifs)
}
