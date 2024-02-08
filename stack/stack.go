package stack

import "iter"

// Stack represents a simple LIFO stack.
type Stack[T any] struct {
	items []T
}

func New[T any](items ...T) *Stack[T] {
	return &Stack[T]{items: items}
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}

	popped := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]

	return popped, true
}

func (s *Stack[T]) TopDown() iter.Seq[T] {
	return func(yield func(T) bool) {
		for i := len(s.items) - 1; i >= 0; i-- {
			if !yield(s.items[i]) {
				return
			}
		}
	}
}

func (s *Stack[T]) BottomUp() iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, item := range s.items {
			if !yield(item) {
				return
			}
		}
	}
}
