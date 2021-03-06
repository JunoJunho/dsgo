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

func last(head *node) *node {
	var target *node = nil
	for iter := head; iter.next != nil; iter = iter.next {
		target = iter.next
	}
	return target
}

func addFirst(head *node, target *node) *node {
	head.prev = target
	target.next = head

	return target // New head
}

func addLast(head *node, target *node) {

	var last_node = last(head)
	last_node.next = target

}

func removeFirst(head *node) *node {
	var second = head.next
	second.prev = nil
	return second
}

func removeLast(head *node) {
	var before_node *node = nil
	for iter := head; iter.next != nil; iter = iter.next {
		before_node = iter
	}
	before_node.next = nil
}
