package single_linked_list

type Node struct {
	data int
	next *Node
}

func Size(first Node) int {
	_size := 1
	for iter := first; iter.next != nil; iter = *(iter.next) {
		_size++
	}
	return _size
}

func last(first Node) Node {
	var _ret Node = first
	for iter := first; iter.next != nil; iter = *iter.next {
		_ret = *iter.next
	}
	return _ret
}

func addFirst(first Node, target Node) Node {
	var _first Node = first
	target.next = &_first // Now target is the first of the Node!

	return target
}

func addLast(first Node, target Node) {
	var _last *Node = nil
	for iter := first; iter.next != nil; iter = *iter.next {
		_last = iter.next
	}
	_last.next = &target
}

func removeFirst(first Node) Node {
	var _ret Node
	return _ret
}
