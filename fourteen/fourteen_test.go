package fourteen

// useful range is restricted  errs out at around 5e+9
// why did they not use 16 bits?  Could probably use same process to extend by two bits
// need to re-examine the decision on whether exp or mant should have bits added
// Use case?

import (
	"fmt"
	"math"
	"testing"
)

//
// TESTS
//

var referenceTests = []struct {
	f14 uint64
	f64 float64
}{
	// special vaues
	{0x0000, pZeroFloat64},       // 00 0000 0000 0000 +0
	{0x2000, nZeroFloat64},       // 10 0000 0000 0000 -0
	{0x1F80, qNaNLowestFloat64},  // 01 1111 1000 0000 QNaN set (low)
	{0x1FFF, qNaNHighestFloat64}, // 01 1111 1111 1111 QNaN set (high)
	{0x1F01, sNaNLowestFloat64},  // 01 1111 0000 0001 SNaN clear (low)
	{0x1F7F, sNaNHighestFloat64}, // 01 1111 0111 1111 SNaN clear (high)
	{0x1F00, pInfinityFloat64},   // 01 1111 0000 0000 +inf
	{0x3F00, nInfinityFloat64},   // 11 1111 0000 0000 -inf

	{0x0F80, +1.0},    // 00 1111 1000 0000 +1
	{0x2F80, -1.0},    // 10 1111 1000 0000 -1
	{0x1120, +10.0},   // 01 0001 0010 0000 +10
	{0x3120, -10.0},   // 11 0001 0010 0000 -10
	{0x147A, +1000.0}, // 01 0100 0111 1010 +1000
	{0x347A, -1000.0}, // 11 0100 0111 1010 -1000
	{0x183C, 192512},  // 01 1000 0011 1100 192512

	// Unrepresentable "ideal" values
	// {0x17C3, +100000.0}, // 01 0111 1100 0011 +100000
	// {0x37C3, -100000.0}, // 01 0111 1100 0011 -100000
	// {0x0FB5, +1.415},    // 00 1111 1011 0101 +1.415
	// {0x2FB5, -1.415},    // 10 1111 1011 0101 -1.415
	// {0x0D5D, +0.054},    // 00 1101 0101 1101 +0.054
	// {0x2D5D, -0.054},    // 10 1101 0101 1101 -0.054
	// {0x0A60, 0.000856}   // 00 1010 0110 0000 0.000856

	// Closest approximations to ideal values
	{0x17C3, +99840.0},        // 01 0111 1100 0011 +100000
	{0x37C3, -99840.0},        // 01 0111 1100 0011 -100000
	{0x12C9, +100.5},          // 01 0010 1100 1001 +100.5
	{0x32C9, -100.5},          // 01 0010 1100 1001 -100.5
	{0x0FB5, +1.4140625},      // 00 1111 1011 0101 +1.415
	{0x2FB5, -1.4140625},      // 10 1111 1011 0101 -1.415
	{0x0D5D, +0.053955078125}, // 00 1101 0101 1101 +0.054
	{0x2D5D, -0.053955078125}, // 10 1101 0101 1101 -0.054
	{0x0A60, 0.0008544921875}, // 00 1010 0110 0000 0.000856
}

// Convert a small test suite of 14-bit floating point values
func TestReference(t *testing.T) {
	for i, a := range referenceTests {
		f64 := FromF14(a.f14)
		f14 := ToF14(a.f64)
		always := false
		if always || f64 != a.f64 && (f64 == f64 || a.f64 == a.f64) || a.f14 != f14 {
			t.Errorf("#%d, decoding 0x%04x is %v; want %v", i, a.f14, f64, a.f64)
			if true {
				t.Errorf("        f14 = 2r%014b", a.f14)
				t.Errorf("        f64 = 2r%064b", math.Float64bits(f64))
				t.Errorf("       want = 2r%064b", math.Float64bits(a.f64))
				t.Errorf("       back = 2r%014b", f14)
			}
		}
	}
}

// Count number of symmetric decode/encode pairs in domain of 14-bit floating point
func TestSymmetry(t *testing.T) {
	match := 0
	count := 0
	for v := uint64(0); v < 1<<14; v++ {
		f64 := FromF14(v)
		f14 := ToF14(f64)
		//fmt.Printf("%5v\n",f14)
		if v == f14 {
			match++
		} else {
			t.Errorf("asymmetric pair: %5d (0x%04x) => %v => %v (0x%04x)", v, v, f64, f14, f14)
		}
		count++
	}
	//fmt.Printf("count = %d\n",count)

	if match != count {
		t.Errorf("total: %d/%d (%.4f%%) match, should be 100%% ", match, count, 100.0*float64(match)/float64(count))
	}
}

//
// EXAMPLES
//

func ExampleFromF14() {
	fmt.Println(FromF14(0x1228)) // the galactic Question?
	// Output:
	//42
}

func ExampleToF14() {
	fmt.Printf("0x%04x\n", ToF14(42))
	// Output:
	//0x1228
}

//
// BENCHMARKS
//

// Measure the average time it takes to convert a 14-bit floating point values

func BenchmarkFromF14(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = FromF14(0x0FB5) // 00 1111 1011 0101 +1.415 (1.4140625)
	}
}

func BenchmarkToF14(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = ToF14(1.4140625) // 00 1111 1011 0101 +1.415 (1.4140625)
	}
}

func TestMdr(t *testing.T) {
	// starting value final value
	for sv := 0.598765; sv < 1e+20; sv *= 10.0 {
		f14 := ToF14(sv)
		fv := FromF14(f14)
		if fv < 0 {
			break
		}
		fmt.Printf("result of From14(ToF14(%g)) = %v\n", sv, fv)
	}
}
