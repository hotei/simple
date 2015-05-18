// pause.go (c) David Rook 2010 - released under Simplified BSD 2-clause License

package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func andy_pause() {
	fmt.Printf("andy_pause:Hit a key to continue\n")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func mdr_pause() {
	fmt.Printf("mdr_pause:Hit a key to continue\n")
	var keybuf [1]byte
	_, _ = os.Stdin.Read(keybuf[0:1])
}

func pause4breakkey(breakKey byte) {
	var keybuf [2]byte

	for {
		// do something interesting here... foo()
		fmt.Printf("pause4breakkey: Any key to keep looping or q to quit\n")
		n, _ := os.Stdin.Read(keybuf[0:2])
		if keybuf[0] == breakKey {
			break
		}
		fmt.Printf("Read in %d chars\n", n)
	}
	fmt.Println()
}

func pause4anykey() {
	// set raw mode for tty
	cmd := exec.Command("/bin/stty", "-F", "/dev/tty", "-icanon", "min", "1") // worked ok
	err := cmd.Run()
	if err != nil {
		fmt.Printf("pause4anykey failed\n")
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("pause4anykey: Hit any key to quit\n")
	var keybuf [1]byte
	_, _ = os.Stdin.Read(keybuf[0:1])
	fmt.Println()
	// reset tty to cooked mode?
	/*
		// not working correctly
		func restore_tty() {
			var argv []string = []string{"/bin/stty", "-F", "/dev/tty", "sane"}
			cmd, err := exec.Run("stty", argv, nil, "/bin/", exec.PassThrough, exec.PassThrough, exec.PassThrough)
			cmd = cmd
			if err != nil {
				fmt.Printf("Cant set sane mode for tty\n")
				os.Exit(1)
			}
		}
	*/
}

func main() {

	fmt.Printf("<Start pause.go>\n")
	fmt.Printf("Andy Pause\n")
	andy_pause()
	fmt.Printf("mdr Pause\n")
	mdr_pause()
	pause4breakkey('q')
	fmt.Printf("Pause4breakkey\n")
	pause4anykey()
	fmt.Printf("Pause4anykey\n")
	fmt.Printf("<End pause.go>\n")

}
