package utils

import "container/heap"

// A queue, used for finding the shortest path
// so we pull the lowest cost item off first

type Item struct {
	value    any // Whatever we are storing in the queue
	priority int // The priority of the item in the queue
	index    int // The index of the item in the heap
}

func (item Item) GetValue() any {
	return item.value
}

func (item Item) GetPriority() int {
	return item.priority
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

// Less converts priority to order, so if the priority is lower,
// we want to be at the start
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	// We've swapped the items, so we need to update the index
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(item any) {
	cast_item := item.(*Item)
	cast_item.index = len(*pq)
	*pq = append(*pq, cast_item)
}

// Grab the item from the end of the list
func (pq *PriorityQueue) Pop() any {
	old_queue := *pq
	len := len(old_queue)
	item := old_queue[len-1]

	// nil the last item, so it can be garbage collected
	old_queue[len-1] = nil

	*pq = old_queue[0 : len-1]
	return item
}

func (pq *PriorityQueue) Update(item *Item, priority int) {
	item.priority = priority
	heap.Fix(pq, item.index)
}

func NewItem(value any, priority int, index int) *Item {
	return &Item{
		value:    value,
		priority: priority,
		index:    index,
	}
}
