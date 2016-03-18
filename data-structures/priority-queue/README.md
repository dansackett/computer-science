# Priority Queue

A priority queue is an abstract data structure which differs from a queue
based on how the items are extracted. Instead of a FIFO or LIFO structure, PQs
return the items based on a set priority. Depending on the underlying
mechanism a PQ can pull the largest or smallest items off first. This kind of
process enables a fast way of extracting minimum and maximum values of a
container.

A heap is one of the most used data structures for delivering the results and
either a min or max heap can be used depending on how the priority will be
set. Another common way of implementing it is with a doubly linked list which
provides some more efficient time complexities on insert and delete in most
cases.

## Use Cases

A priority queue, like queues and stacks, can be used in a lot of ways. The
difference in applications deals more with the fact that LIFO and FIFO don't
apply here. It's all about an item's priority. Some use cases could be:

- Bandwidth management: By using a PQ here, you can ensure the proper traffic
  dispersion is maintained. For instance, real-time traffic will have a higher
  priority.
- Dijkstra's Algorithm: A PQ allows extracting a minimum vertex efficiently.
- Huffman Coding: Since you need to find the two lowest-frequency trees then a
  PQ offers that kind of functionality.
- Job Scheduling: For jobs that need to occur before others, a PQ offers a
  sane way to manage job hierarchy.

## Time Complexity

Since the PQ is using a Max Heap as the underlying mechanism for operations,
the time complexity is that of a binary max heap.

| Case      | Average       | Worst Case    |
| --------- |:-------------:|:-------------:|
| Space     | `O(n)`        | `O(n)`        |
| Search    | `O(n)`        | `O(n)`        |
| Insert    | `O(lg n)`     | `O(lg n)`     |
| Delete    | `O(lg n)`     | `O(lg n)`     |
