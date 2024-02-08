package stack

import "iter"

// Stack represents a simple LIFO stack.
type Stack[T any] struct {
	items []T
}

// New returns a stack hydrated with the given items.
func New[T any](items ...T) *Stack[T] {
	return &Stack[T]{items: items}
}

// Push adds an item to the top of the stack.
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// Pop removes and returns the item at the top of the stack.
func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}

	popped := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]

	return popped, true
}

// TopDown returns an iterator that walks down the stack from the top.
func (s *Stack[T]) TopDown() iter.Seq[T] {
	return func(yield func(T) bool) {
		for i := len(s.items) - 1; i >= 0; i-- {
			if !yield(s.items[i]) {
				return
			}
		}
	}
}

// TopDown returns an iterator that walks up the stack from the bottom.
func (s *Stack[T]) BottomUp() iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, item := range s.items {
			if !yield(item) {
				return
			}
		}
	}
}
