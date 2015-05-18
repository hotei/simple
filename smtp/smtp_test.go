// smtp_test.go

// go test [-bench=".*"]   # re2 expression matches everything, runs all benchmarks

package x

import (
	"bytes"
	"fmt"
	"github.com/hotei/mdr"
	"io"
	"io/ioutil"
	"testing"
)

// template
func Test_000(t *testing.T) {
	fmt.Printf("Test_000 \n")
	if false {
		t.Errorf("print fail, but keep testing")
	}
	if false {
		t.Fatalf("print fail and keep testing")
	}
	fmt.Printf("go test -bench=\".*\" to run all benchmarks\n")
	fmt.Printf("running AndyMain\n")
	AndyMain()
	fmt.Printf("done with AndyMain\n")
	fmt.Printf("Pass - test 000\n")
}

/////////////////////////  B E N C H M A R K S  ////////////////////////////
/*  4 GHz AMD-64  8120 8 core
Benchmark_PseudoRandomBlock-8	      50	  36,471,550 ns/op (AMD-64 4000)
Benchmark_PseudoRandomBlock-8	      50	  20,355,419 ns/op (i7-2500) go 1.4.2
*/

// 46.9e6 ns/op on 4Ghz AMD64 with 1.0.3
// 36.6e6 ns/op on 4Ghz AMD64 with 1.1 << 22% better >>
// 35.1e6 ns/op on 4Ghz AMD64 with 1.2
func Benchmark_PseudoRandomBlock(b *testing.B) {
	PRBsize := 1000000
	for i := 0; i < b.N; i++ {
		x := mdr.PseudoRandomBlock(PRBsize)
		r := bytes.NewReader(x)
		if _, err := io.Copy(ioutil.Discard, r); err != nil {
			panic(err)
		}
	}
}
