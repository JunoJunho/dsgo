package stack

import "github.com/JunoJunho/dsgo/single_linked_list"

type stack struct {
	floor *single_linked_list.Node // Always floor?
	// Grow from floor to top
}

func isEmpty(s *stack) bool {
	if s.floor == nil {
		return true
	}
	return false
}

func push(s *stack, data int) {
	var new_node = single_linked_list.Node{Data: data, Next: nil}
	single_linked_list.AddLast(*s.floor, new_node)
}
