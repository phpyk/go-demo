package simplemath

import (
	"testing"
)

func TestAdd(t *testing.T) {
	r := Add(1, 2)
	if r != 3 {
		t.Errorf("Add(1,2) faild, Got %d, excepted 3.", r)
	}
}

