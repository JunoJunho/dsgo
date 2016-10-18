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
		t.Errorf("The last must be 2! not %d", _ret.data)
	}
}

func TestAddFirst(t *testing.T) {
	var head = node{0, nil, nil}
	var second = node{1, nil, nil}
	var third = node{2, nil, nil}

	head.next = &second
	second.prev = &head

	var new_head = addFirst(&head, &third)
	if new_head.data != 2 {
		t.Errorf("The first data should be 2! not %d", new_head.data)
	}
}

func TestAddLast(t *testing.T) {
	var head = node{0, nil, nil}
	var second = node{1, nil, nil}
	var third = node{2, nil, nil}

	head.next = &second
	second.prev = &head

	addLast(&head, &third)
	var last_node = last(&head)
	if last_node.data != 2 {
		t.Errorf("The first data should be 2! not %d", last_node.data)
	}
}

func TestRemoveFirst(t *testing.T) {
	var head = node{0, nil, nil}
	var second = node{1, nil, nil}
	head.next = &second
	second.prev = &head

	var new_head = removeFirst(&head)

	if new_head.data != 1 {
		t.Errorf("The first data should be 1! not %d", new_head.data)
	}
}
