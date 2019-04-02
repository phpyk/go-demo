package simplemath

import (
	"testing"
)

func TestSqrt(t *testing.T) {
	v := Sqrt(16)
	if v != 4 {
		t.Errorf("Sqrt(16) faild. Got %v, excepted 4.", v)
	}
}
