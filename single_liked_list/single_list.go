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

//
// func addFirst(first node, target node) {
//
// }
//
// func addLast(first node, target node) {
//
// }

func removeFirst(first node) node {
	var _ret node
	return _ret
}
