package heap

// MinHeap is a wrapper around a heap and will work the minimum values coming
// off first when popping items.
type MinHeap struct {
	*Heap
}

// NewMinHeap returns a new min heap structure.
func NewMinHeap() *MinHeap {
	return &MinHeap{NewHeap()}
}

// Less determines how to compare two keys when heapifying.
func (h *MinHeap) Less(i, j int) bool {
	return h.Heap.items[i].Key < h.Heap.items[j].Key
}

// Push is a wrapper around the heap's main push method which specifies the Up
// function for the callback.
func (h *MinHeap) Push(k int, v interface{}) {
	h.Heap.push(k, v, h.Less)
}

// Pop is a wrapper around the heap's main pop method which specifies the Down
// function for the callback.
func (h *MinHeap) Pop() interface{} {
	return h.Heap.pop(h.Less)
}

// Iter is a wrapper around the heap's Iter method for easy iteration.
func (h *MinHeap) Iter() chan interface{} {
	return h.Heap.Iter()
}

// Peek returns the next item to be popped off the heap.
func (h *MinHeap) Peek() interface{} {
	return h.Heap.peek()
}
