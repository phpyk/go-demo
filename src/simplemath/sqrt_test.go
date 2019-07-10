package simplemath

import (
	"testing"
)

func TestSqrt(t *testing.T) {
	r := Sqrt(4)
	if r != 2 {
		t.Errorf("Sqrt(4) faild Got %d expected: 2",r)
	}
}
