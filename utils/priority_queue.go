package utils

// Implementing a queue sorted in the order
// definied by a comparison function

type PriorityQueue[T comparable] struct {
	items      []T
	comparator func(a, b T) int
}

// NewPriorityQueue creates a new prio queue with the given comparitor
func NewPriorityQueue[T comparable](comparator func(a, b T) int) PriorityQueue[T] {
	return PriorityQueue[T]{comparator: comparator}
}

// Reference https://maneeshaindrachapa.medium.com/heap-data-structure-in-golang-98641a32d2e3

// maxHeapifyDown process
// func (pq *PriorityQueue[T]) maxHeapifyDown(index int) {
// 	lastIndex := len(pq.items) - 1
// 	l, r := left(index), right(index)
// 	childToCompare := 0
// 	// loop while index has at least one child
// 	for l <= lastIndex { // when left child is the only child
// 		if l == lastIndex {
// 			childToCompare = l
// 		} else if pq.items[l] > pq.items[r] { // when left child is larger
// 			childToCompare = l
// 		} else { // when right child is larger
// 			childToCompare = r
// 		}
// 		// Compare array value of the current index to larger child and swap if smaller
// 		if pq.items[index] < pq.items[childToCompare] {
// 			pq.swap((index), childToCompare)
// 			index = childToCompare
// 			l, r = left(index), right(index)
// 		} else {
// 			return
// 		}
// 	}
// }

// Get the left child index
func left(i int) int {
	return 2*i + 1
}

// Get the right child index
func right(i int) int {
	return 2*i + 2
}
