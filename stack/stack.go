package stack

import (
	"fmt"
)

// Stack implemented with the help of slices
type Stack struct {
	values []interface{}
	top    int
}

// NewStack method creates the stack
func NewStack() *Stack {
	s := &Stack{
		top:    -1,
		values: make([]interface{}, 0),
	}
	return s
}

// Push method inserts the item into the stack
func (s *Stack) Push(item interface{}) {
	s.values = append(s.values, item)
	s.top++
}

// Pop method returns the item from the stack
func (s *Stack) Pop() (interface{}, error) {
	var item interface{}

	if s.top >= 0 {
		item = s.values[s.top]
		s.values = append(s.values[:s.top], s.values[s.top+1:]...)
		s.top--
	} else {
		return nil, fmt.Errorf("[ ERR ] Index out of range")
	}

	return item, nil
}

// Size method returns the size of the stack
func (s *Stack) Size() int {
	return (s.top + 1)
}
