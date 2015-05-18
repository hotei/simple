// md5.go (c) David Rook 2010 - released under Simplified BSD 2-clause License

package main

import (
	"crypto/md5"
	"fmt"
	"hash"
	"io"
	"io/ioutil"
	"os"
)

func fatal_err(err error) {
	fmt.Printf("!Err-> %s\n", err)
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
	digest := h.Sum(nil)
	fmt.Printf("%x: %s md5 as computed\n", digest, original)
}

func BytesDigestMd5(bufr []byte) []byte {
	var h hash.Hash = md5.New()
	h.Write(bufr)
	return h.Sum(nil)
}

func FileDigestMd5(fname string) []byte {
	megabuf := make([]byte, 1024*1024)

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
	original := "Memories.mp3"
	digest := FileDigestMd5(original)

	osVersion := "00de9875e48ab9eb77bfc33e28ed806a"
	fmt.Printf("%x Md5 go() of %s\n", digest, original)
	fmt.Printf("%s Md5 from md5sum \n", osVersion)
	if osVersion == fmt.Sprintf("%x", digest) {
		fmt.Printf("match\n")
	} else {
		fmt.Printf("!Err-> do not match")
	}
}
