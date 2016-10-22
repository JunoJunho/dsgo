package stack

import "github.com/JunoJunho/dsgo/single_linked_list"

type stack struct {
	floor *single_linked_list.Node
}

func isEmpty(s *stack) bool {
	if s.floor == nil {
		return true
	}
	return false
}
