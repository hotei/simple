// md5File.go (c) David Rook 2010 - released under Simplified BSD 2-clause License

// +build ignore

package main

import (
	"crypto/md5"
	"fmt"
	"hash"
	"io"
	"io/ioutil"
	"os"
)

const MBUFSIZE = 1024 * 1024

func fatal_err(erx error) {
	fmt.Printf("%s \n", erx)
	os.Exit(1)
}

// reads entire file at once
func slurp() {
	original := "Memories.mp3"
	origbytes, err := ioutil.ReadFile(original)
	if err != nil {
		fatal_err(err)
	}
	var h hash.Hash = md5.New()
	h.Write(origbytes)
	fmt.Printf("%x: %s computed\n", h.Sum(nil), original)

}

func getFileMD5(fname string) []byte {
	fmt.Printf("start getFileMD5 on %s\n", fname)
	megabuf := make([]byte, MBUFSIZE)

	f, err := os.Open(fname)
	if err != nil {
		fatal_err(err)
	}
	defer f.Close()

	var h hash.Hash = md5.New()
	for {
		n, err := io.ReadFull(f, megabuf)
		//        fmt.Printf("Read %d bytes, %v err, %v\n", n,err,os.EOF)
		if err == io.ErrUnexpectedEOF {
			if n == 0 {
				break
			}
			smallbuf := megabuf[0:n]
			h.Write(smallbuf)
			//            fmt.Printf("%x: computed on final piece\n", h.Sum() )
			break
		}
		if n == 0 {
			break
		} // might happen but rare, err will be os.EOF
		if err != nil {
			if err != io.ErrUnexpectedEOF {
				fatal_err(err)
			}
		}
		h.Write(megabuf)
		//        fmt.Printf("%x: computed on 1MB segment\n", h.Sum() )
	}
	//    fmt.Printf("%x: computed\n", h.Sum() )
	return h.Sum(nil)
}

func main() {
	fmt.Printf("<start md5File.go>\n")
	// '00de9875e48ab9eb77bfc33e28ed806a  Memories.mp3 (actual)'

	original := "kpx.iso"
	filehash := getFileMD5(original)
	fmt.Printf("md5 hash = %x\n", filehash)
}
