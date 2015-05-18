// file_test.go (c) David Rook 2010 - released under Simplified BSD 2-clause License

/*
 NOTE there are no t.Errorf calls here.  If any function calls os.Exit()
 the test will fail.
*/

package main

import (
	"fmt"
	"testing"
)

func Test_File(t *testing.T) {
	fmt.Printf("<SimpleFile.go starting>\n")

	simpleReader("./SimpleFile.go")
	simpleWriter("./testfile.txt")
	fmt.Printf("Writing out testfile.txt\n")
	fmt.Printf("Reading in our new testfile.txt\n")
	simpleReader("./testfile.txt")
	fmt.Printf("<SimpleFile.go finished>\n")
}
