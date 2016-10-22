package stack

import "testing"

func TestIsEmpty(t *testing.T) {
	var s = stack{nil}
	if !isEmpty(&s) {
		t.Errorf("Empty stack should return true!")
	}
}
