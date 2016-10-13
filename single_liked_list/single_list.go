package single_list

type node struct {
	data int
	next *node
}

func size(first node) int {
	_size := 1
	for iter := first; iter.next != nil; iter = *(iter.next) {
		_size++
	}
	return _size
}

func last(first node) node {
	var _ret node = first
	for iter := first; iter.next != nil; iter = *iter.next {
		_ret = *iter.next
	}
	return _ret
}

func addFirst(first node, target node) node {
	var _first node = first
	target.next = &_first // Now target is the first of the node!

	return target
}

func addLast(first node, target node) {
	var _last *node = nil
	for iter := first; iter.next != nil; iter = *iter.next {
		_last = iter.next
	}
	_last.next = &target
}

func removeFirst(first node) node {
	var _ret node
	return _ret
}
