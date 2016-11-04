package stack

import "testing"

func TestIsEmpty(t *testing.T) {
	var s = stack{nil}
	if !isEmpty(&s) {
		t.Errorf("Empty stack should return true!")
	}
}

func TestPush(t *testing.T) {
	var s = stack{nil}
	var nd = 3
	push(&s, nd)

	if isEmpty(&s) {
		t.Errorf("After push, stack should not empty!")
	}

}
