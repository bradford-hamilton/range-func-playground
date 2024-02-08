package main

import (
	"fmt"
	"iter"
)

// Queue represents a simple FIFO queue.
type Queue[T any] struct {
	items []T
}

// NewQueue creates a new Queue.
func NewQueue[T any](items ...T) *Queue[T] {
	return &Queue[T]{items: items}
}

// Enqueue adds an item to the queue.
func (q *Queue[T]) Enqueue(item T) {
	q.items = append(q.items, item)
}

// Dequeue removes and returns the item in the front of the
// queue as well as a success flag in case of failure.
func (q *Queue[T]) Dequeue() (T, bool) {
	if len(q.items) == 0 {
		var zeroVal T
		return zeroVal, false
	}

	item := q.items[0]
	q.items = q.items[1:]

	return item, true
}

// All returns an iterator over the queue's items.
func (q *Queue[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, item := range q.items {
			if !yield(item) {
				return
			}
		}
	}
}

// GOEXPERIMENT=rangefunc
func main() {
	queue := NewQueue(1, 2, 3, 4, 5)
	for item := range queue.All() {
		fmt.Println(item)
	}
}
