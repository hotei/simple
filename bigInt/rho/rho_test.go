// rho_test.go
// go test -cpuprofile=rho.prof -bench="Both" -benchmem=true

package rho

import (
	//"fmt"
	"math/big"
	"testing"
)

var N5 *big.Int = FromString("766150476015982127183457373")
var N6 *big.Int = FromString("62823675885202669644104829577")

func BenchmarkBoth(t *testing.B) {
	Rho(N5)
	Rho(N6)
}

//func ExampleN5() {
//	r5 := Rho(N5)
//	fmt.Printf("%v",r5)
//	// Output: 1178524040059
//}

//func ExampleN6() {
//	r6 := Rho(N6)
//	fmt.Printf("%v",r6)
//	// Output: 663410067979
//}
