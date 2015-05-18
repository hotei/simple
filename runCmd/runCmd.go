// runCmd.go (c) David Rook 2010 - released under Simplified BSD 2-clause License

package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func FatalError(err error) {
	fmt.Printf("!Err-> %s\n", err)
	os.Exit(1)
}

func runCmd() error {

	// working
	fmt.Printf("Use tr to uppercase a string\n")
	cmd := exec.Command("tr", "a-z", "A-Z")
	cmd.Stdin = strings.NewReader("Some Input")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("in all caps: %q\n", out.String())

	// working
	fmt.Printf("executing command: ls -l /etc/fstab\n")
	cmd = exec.Command("ls", "-l", "/etc/fstab")
	cmd.Stdin = strings.NewReader("")
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		FatalError(err)
	}
	fmt.Printf("%s\n", out.String())
	fmt.Printf("<done>\n")
	return nil
}

func main() {
	if err := runCmd(); err != nil {
		FatalError(err)
	}
}
