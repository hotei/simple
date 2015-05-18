// fourteen.go  - just copied from go-nuts thread - none of this is mine

package fourteen

/*
	14-bit floating point conversion functions

	Here's a sample calculation by N. Riesco (nicolas.riesco@gmail.com) for a
	float14 (1 sign bit, 6 bits to encode the exponent and 7 bits to encode
	the mantissa):

	0 | 100101 | 1001001

	sign = 0
	exp + bias = 0b100101 = 37
	mantissa without leading 1 = 0b1001001

	exp = 37 - bias = 37 - 31 = 6
	mantissa = 0b1.1001001 = 0b11001001 / 2^7 = 1.5703125

	100.5 = mantissa * 2 ^ exponent = 1.5703125 * 2^6
*/

import "math"

//
// Special Values (http://www.doc.ic.ac.uk/~eedwards/compsys/float/nan.html)
//

// positive zero
var pZeroUint14 = uint64(0)
var pZeroUint64 = uint64(0)
var pZeroFloat64 = math.Float64frombits(pZeroUint64)

// negative zero
var nZeroUint14 = uint64(1 << 13)
var nZeroUint64 = uint64(1 << 63)
var nZeroFloat64 = math.Float64frombits(nZeroUint64)

// positive infinity
var pInfinityUint14 = uint64(0x1F00)
var pInfinityFloat64 = math.Inf(+1)
var pInfinityUint64 = math.Float64bits(pInfinityFloat64)

// negative infinity
var nInfinityUint14 = uint64(0x3F00)
var nInfinityFloat64 = math.Inf(-1)
var nInfinityUint64 = math.Float64bits(nInfinityFloat64)

// signaling NaN with lowest payload value
var sNaNLowestUint14 = uint64(0x1F01) // 01 1111 0000 0001 SNaN clear (low)
var sNaNLowestUint64 = uint64(0x7FF0000000000001)
var sNaNLowestFloat64 = math.Float64frombits(sNaNLowestUint64)

// signaling NaN with highest payload value
var sNaNHighestUint14 = uint64(0x1F7F) // 01 1111 0111 1111 SNaN clear (high)
var sNaNHighestUint64 = uint64(0x7FF7FFFFFFFFFFFF)
var sNaNHighestFloat64 = math.Float64frombits(sNaNHighestUint64)

// quiet NaN with lowest payload value
var qNaNLowestUint14 = uint64(0x1F80) // 01 1111 1000 0000 QNaN set (low)
var qNaNLowestUint64 = uint64(0x7FF8000000000000)
var qNaNLowestFloat64 = math.Float64frombits(qNaNLowestUint64)

// quiet NaN with highest payload value
var qNaNHighestUint14 = uint64(0x1FFF) // 01 1111 1111 1111 QNaN set (high)
var qNaNHighestUint64 = uint64(0x7FFFFFFFFFFFFFFF)
var qNaNHighestFloat64 = math.Float64frombits(qNaNHighestUint64)

//
// Conversions
//

func FromF14(v uint64) float64 {
	// function call costs 2.51 ns/op with immediate return (2.7 GHz intel Core i7 / MacPro)
	//return 0

	// function call for special cases costs 4.6 ns
	switch v {
	case pZeroUint14: // 0x0000
		return pZeroFloat64
	case nZeroUint14: // 0x2000
		return nZeroFloat64
	case pInfinityUint14: // 0x1F00
		return pInfinityFloat64
	case nInfinityUint14: // 0x3F00
		return nInfinityFloat64
	case sNaNLowestUint14: // 0x1F01
		return sNaNLowestFloat64
	case sNaNHighestUint14: // 0x1F7F
		return sNaNHighestFloat64
	case qNaNLowestUint14: // 0x1F80
		return qNaNLowestFloat64
	case qNaNHighestUint14: // 0x1FFF
		return qNaNHighestFloat64
	}

	// function call for conversion costs 3.21 ns/op as step-by-step extraction (5.31 ns/op total)
	// s := (v >> 13) & 0x01                 // sign
	// e := ((v >> 7) & 0x3f) + 0x3ff - 0x1f // biased exponent as stored
	// f := v & 0x7f                         // fraction
	// return math.Float64frombits((s << 63) | (e << 52) | (f << 45))

	// function call for conversion costs 2.62 ns/op as simple expression (5.01 ns/op total)
	return math.Float64frombits((((v & 0x2000) << 50) | ((v & 0x1fff) << 45)) + ((0x3ff - 0x1f) << 52))
}

func ToF14(v float64) uint64 {
	b := math.Float64bits(v)

	switch b {
	case pZeroUint64:
		return pZeroUint14
	case nZeroUint64:
		return nZeroUint14
	case pInfinityUint64:
		return pInfinityUint14
	case nInfinityUint64:
		return nInfinityUint14
	case sNaNLowestUint64:
		return sNaNLowestUint14
	case sNaNHighestUint64:
		return sNaNHighestUint14
	case qNaNLowestUint64:
		return qNaNLowestUint14
	case qNaNHighestUint64:
		return qNaNHighestUint14
	}

	// NOTE: We do not check the exponent for range, therefore some IEEE 754 floats
	// will silently exceed the range of the smaller exponent. This is fine for the
	// intended use, which is decoding existing 14-bit values or encoding values in
	// the proper domain. We also don't check for underflow of any denormal values
	// during the fraction processing. If desired, add checks for "e > 0x7f" before
	// the final AND and also "(b >> 45) != 0 && b == 0" after the AND.

	s := (b >> 63) & 0x01                  // sign
	e := ((b >> 52) - 0x3ff + 0x1f) & 0x7f // biased exponent as stored
	f := (b >> 45) & 0x7f                  // fraction
	return (s << 13) | (e << 7) | f
}
