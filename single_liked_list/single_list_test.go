package single_list

import "testing"

func TestSingleLinkedList(t *testing.T) {
	var first = node{1, nil}
	var second = node{2, nil}
	var third = node{3, nil}
	first.next = &second
	second.next = &third

	if size(first) != 3 {
		t.Errorf("Size is 3!! not %d", size(first))
	}
}

func TestLast(t *testing.T) {
	var first = node{1, nil}
	var second = node{2, nil}
	var third = node{3, nil}
	first.next = &second
	second.next = &third
	ret := last(first)
	if ret.data != 2 {
		t.Errorf("Last is 3!! not %d", ret.data)
	}

}
