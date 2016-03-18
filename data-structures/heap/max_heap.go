package heap

// MaxHeap is a wrapper around a heap and will work the maximum values coming
// off first when popping items.
type MaxHeap struct {
	*Heap
}

// NewMaxHeap returns a new max heap structure.
func NewMaxHeap() *MaxHeap {
	return &MaxHeap{NewHeap()}
}

// Less determines how to compare two keys when heapifying.
func (h *MaxHeap) Less(i, j int) bool {
	return h.Heap.items[i].Key > h.Heap.items[j].Key
}

// Push is a wrapper around the heap's main push method which specifies the Up
// function for the callback.
func (h *MaxHeap) Push(k int, v interface{}) {
	h.Heap.push(k, v, h.Less)
}

// Pop is a wrapper around the heap's main pop method which specifies the Down
// function for the callback.
func (h *MaxHeap) Pop() interface{} {
	return h.Heap.pop(h.Less)
}

// Iter is a wrapper around the heap's Iter method for easy iteration.
func (h *MaxHeap) Iter() chan interface{} {
	return h.Heap.Iter()
}

// Peek returns the next item to be popped off the heap.
func (h *MaxHeap) Peek() interface{} {
	return h.Heap.peek()
}
