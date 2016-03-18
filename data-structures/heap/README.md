# Binary Heap

A heap is a tree-like data structure. I'm going to focus on a binary heap for
now in order to keep it simple. Binary heaps follow the contraints:

- The binary tree is a complete binary tree with all levels full except for
  the last level which is filled from left to right.
- All nodes are either greater than or equal to or less than or equal to each
  of its children.

When talking about a binary heap, it's common to consider both a **Max Heap** and
a **Min Heap**. A max heap ensures that the largest value is at the top of the
heap and can be reached in O(1) time. A min heap is the opposite with the
smallest value at the top of the heap. Both can be implemented the same way
and only differ in the comparison of items when updating the heap.

One of the nice things about heaps is that they can be handled entirely in an
array reducing the memory cost of pointers.

Operations on a binary heap will typically be pushing new items onto the heap
or extracting the minimum (or maximum) value from the heap. In both cases we
do core operations and then update the heap property to account for the new or
now missing values.

For insertion this means appending the new item to the end of the array and
then bubbling upwards checking itself against its parent swapping values until
the root is reached. In this way, it's ensured that the minimum (or maximum)
value always reaches the top of the tree.

For extraction, this means popping the first item from the array and bubbling
downwards based on the smaller (or larger) child node. A swap can be done in
this direction as well making the heap property satisfied again.

One important thing to know when bubbling both up and down is how to locate
the children. Since we know that the tree is complete and that each item has
at most two children then we can assume the following:

```
Assuming i = current index
Left Child = 2 * i
Right Child = 2 * i + 1
Parent = floor(i / 2)
```

With these, we can freely traverse the tree correcting the heap properties.

## Use Cases

The heap can be used in a lot of different ways. Some of the more well known
implementations are:

- Heapsort: Very good in-place sorting algorithm
- Select Algorithms: Cases where you need to find min, max, or kth element are
  more efficient with a heap.
- Graph Algorithms: Can be used as a traversal structure within.
- Priority Queue: One of the bigger use cases of a heap, a priority queue
  is an abstract data-type which allows for attaching priorities to queued
  items.
- Order Statistics: Heaps can help in finding the kth smallest or largest
  elements in an array.

## Time Complexity

| Case      | Average       | Worst Case    |
| --------- |:-------------:|:-------------:|
| Space     | `O(n)`        | `O(n)`        |
| Search    | `O(n)`        | `O(n)`        |
| Insert    | `O(lg n)`     | `O(lg n)`     |
| Delete    | `O(lg n)`     | `O(lg n)`     |
