package stax

import "context"

// Stax is a stack machine interface
type Stax interface {
	Push(item Item)
	Pop() Item
	Peek() Item
}

var _ Stax = (*Stack)(nil)

// Interpreter defines a Run function that runs a Stax.
type Interpreter interface {
	Run(ctx context.Context) error
}

// Item is an element in the stack
type Item struct {
	Instruction string
}

// Stack holds a stack of items
type Stack struct {
	// TODO: protect stack with a mutex (do we need concurrency?)
	stack []Item
}

// NewStack returns a new stack with the given Items
func NewStack(items []Item) Stax {
	return &Stack{
		stack: items,
	}
}

// Push adds an item to the stack
func (s *Stack) Push(item Item) {
	s.stack = append(s.stack, item)
}

// Pop removes the top item of the stack and returns it
func (s *Stack) Pop() Item {
	if len(s.stack) == 0 {
		// TODO: what should we do here? exit program when we add instruction set?
		panic("not impl")
	}
	x, y := s.stack[len(s.stack)-1], s.stack[:len(s.stack)-1]
	s.stack = y
	return x
}

// Peek returns the top of the stack but does not remove it
func (s *Stack) Peek() Item {
	if len(s.stack) == 0 {
		// TODO: what should we do here? exit program when we add instruction set?
		panic("not impl")
	}
	x := s.stack[len(s.stack)-1]
	return x
}
