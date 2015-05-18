// httpd.go (c) David Rook 2010 - released under Simplified BSD 2-clause License

// working - 'hello world' for httpd

package main

import (
	"fmt"
	"log"
	"net/http"
)

func version_1() {
	err := http.ListenAndServe(":12345", http.FileServer(http.Dir("/home/mdr/Desktop/DOC_vwar/")))
	if err != nil {
		log.Printf("error running docs webserver: %v", err)
	}
}

func main() {
	fmt.Printf("Starting http server at %d\n", 12345)
	version_1()
}
