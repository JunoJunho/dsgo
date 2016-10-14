package doubly_linked_list

type node struct {
	data int
	prev *node
	next *node
}

func size(head *node) int {
	var i int = 0
	if head != nil {
		i += 1
	}
	for iter := head; iter.next != nil; iter = iter.next {
		i += 1
	}
	return i
}
