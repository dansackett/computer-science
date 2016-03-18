package heap

// Item is a key / value pair for people who want to set the priority and
// arbitrary data.
type Item struct {
	Key int
	Val interface{}
}

// NewItem returns a new item instance.
func NewItem(key int, val interface{}) *Item {
	return &Item{Key: key, Val: val}
}

// Heap is our actual heap data structure. You can set a key / value pair with
// the key being the priority.
type Heap struct {
	items []*Item
	size  int
}

// NewHeap returns a new Heap instance.
func NewHeap() *Heap {
	var items []*Item
	return &Heap{items: items, size: 0}
}

// Cardinality returns the size of the queue.
func (h *Heap) Cardinality() int {
	return h.size
}

// IsEmpty returns whether the Heap has items in it.
func (h *Heap) IsEmpty() bool {
	return h.Cardinality() == 0
}

// Swap trades items at specified indices.
func (h *Heap) Swap(i, j int) {
	h.items[i], h.items[j] = h.items[j], h.items[i]
}

// peek is the internal function used to view the next item in the heap.
func (h *Heap) peek() interface{} {
	return h.items[0].Val
}

// Iter returns a channel to iterate through the Heap items.
func (h *Heap) Iter() chan interface{} {
	ch := make(chan interface{})
	go func() {
		for _, v := range h.items {
			ch <- v.Val
		}
		close(ch)
	}()
	return ch
}

// push is the internal function called by the min / max heap to add a new
// item to the heap and then fix the heap properties.
func (h *Heap) push(k int, v interface{}, lessFunc func(i, j int) bool) {
	h.items = append(h.items, NewItem(k, v))
	h.size++
	h.up(h.size-1, lessFunc)
}

// pop is the internal function called by the min / max heap to remove the
// smallest item from the heap and then fix the heap properties.
func (h *Heap) pop(lessFunc func(i, j int) bool) interface{} {
	if h.IsEmpty() {
		return -1
	}
	ret := h.items[0]
	h.items[0] = h.items[h.size-1]
	h.size--
	h.items = h.items[1:]
	h.down(0, lessFunc)
	return ret.Val
}

// up moves up the heap from the back of the internal items list resorting
// items as it bubbles up so the heap property is restored.
func (h *Heap) up(i int, lessFunc func(i, j int) bool) {
	parent := (i - 1) / 2
	if i != parent && lessFunc(i, parent) {
		h.Swap(i, parent)
		h.up(parent, lessFunc)
	}
}

// down moves down the heap from the initial index resorting items as it
// bubbles down so the heap property is restored.
func (h *Heap) down(i int, lessFunc func(i, j int) bool) {
	left := i*2 + 1
	right := i*2 + 2
	min := i

	if left < h.size && lessFunc(left, i) {
		min = left
	} else if right < h.size && lessFunc(right, i) {
		min = right
	}

	if i != min {
		h.Swap(i, min)
		h.down(min, lessFunc)
	}
}
