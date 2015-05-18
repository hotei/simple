// funcParm_test.go (c) David Rook 2010 - released under Simplified BSD 2-clause License

package main

import (
	"testing"
)

func Test_01(t *testing.T) {
	if addit(4, 6) != 10 {
		t.Errorf("addit() failed in Test_01()")
	}

	if diffit(10, 6) != 4 {
		t.Errorf("diffit() failed in Test_01()")
	}

	if funtest(3, 9, addit) != 12 {
		t.Errorf("funtest(addit) failed in Test_01()")
	}

	if funtest(12, 9, diffit) != 3 {
		t.Errorf("funtest(diffit) failed in Test_01()")
	}
}
