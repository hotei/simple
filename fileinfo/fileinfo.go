// Simple/fileinfo.go

// +build darwin freebsd linux netbsd openbsd

package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	fmt.Printf("startup\n")
	fi, err := os.Stat("fileinfo.go")
	if err != nil {
		fmt.Printf("err %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("fi %v\n\n\n", fi)
	sys := fi.Sys().(*syscall.Stat_t)
	fmt.Printf("UID = %d\n", int(sys.Uid))
	fmt.Printf("GID = %d\n", int(sys.Gid))
}
