package doubly_linked_list

import "testing"

func TestSize(t *testing.T) {
	var head = node{0, nil, nil}
	var second = node{1, nil, nil}
	var third = node{2, nil, nil}
	var fourth = node{3, nil, nil}

	head.next = &second
	second.prev = &head
	second.next = &third
	third.prev = &second
	third.next = &fourth

	var _ret = size(&head)

	if _ret != 4 {
		t.Errorf("The size must be 4! not %d", _ret)
	}
}

func TestLast(t *testing.T) {
	var head = node{0, nil, nil}
	var second = node{1, nil, nil}
	var third = node{2, nil, nil}

	head.next = &second
	second.prev = &head
	second.next = &third
	third.prev = &second

	var _ret = last(&head)
	if _ret.data != 2 {
		t.Errorf("The size must be 2! not %d", _ret.data)
	}
}
