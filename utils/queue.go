package utils

// Implementing a queue based off of an array list
// This queue is FIFO

type Queue[T any] struct {
	items []T
}

// Push to the top of the queue
func (q *Queue[T]) Push(n T) {
	q.items = append(q.items, n)
}

// Pop from the bottom of the queue
func (q *Queue[T]) Pop() T {
	lastIndex := len(q.items) - 1
	top := q.items[lastIndex]
	q.items = q.items[:lastIndex]
	return top
}

// Peek the top item of a queue
func (q *Queue[T]) Peek() T {
	return q.items[len(q.items)-1]
}

// Get the length of the queue
func (q *Queue[T]) Len() int {
	return len(q.items)
}

// Check if queue is empty
func (q *Queue[T]) IsEmpty() bool {
	return len(q.items) == 0
}
