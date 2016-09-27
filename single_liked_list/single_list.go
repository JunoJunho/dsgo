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
