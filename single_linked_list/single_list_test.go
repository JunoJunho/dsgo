package single_linked_list

import "testing"

func TestSingleLinkedList(t *testing.T) {
	var first = Node{1, nil}
	var second = Node{2, nil}
	var third = Node{3, nil}
	first.next = &second
	second.next = &third

	if Size(first) != 3 {
		t.Errorf("Size is 3!! not %d", Size(first))
	}
}

func TestLast(t *testing.T) {
	var first = Node{1, nil}
	var second = Node{2, nil}
	var third = Node{3, nil}
	first.next = &second
	second.next = &third
	ret := last(first)
	if ret.data != 3 {
		t.Errorf("Last is 3!! not %d", ret.data)
	}
}

func TestAddFirst(t *testing.T) {
	var first = Node{1, nil}
	var second = Node{2, nil}
	var new_first = Node{0, nil}
	first.next = &second

	var new_list = addFirst(first, new_first)

	if new_list.data != 0 {
		t.Errorf("AddFirst must return 0!! not %d", new_list.data)
	}
}

func TestAddLast(t *testing.T) {
	var first = Node{1, nil}
	var second = Node{2, nil}
	var new_last = Node{3, nil}
	first.next = &second

	addLast(first, new_last)
	var last_data = last(first)
	if last_data.data != 3 {
		t.Errorf("Last is 3!! not %d", last_data.data)
	}
}
