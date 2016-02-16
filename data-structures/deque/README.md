# Deque

A deque or a "double-ended queue" is an abstact data type which works much
like a queue except we can perform operations on either end of the queue.
These operations include appending, prepending, popping, and popping left.

In this implementation, I'm using a circular doubly linked list as the
container.

## Use Cases

A deque is funny in that it doesn't have a concrete usage as a normal queue or
stack can be used in a lot of the places where this is needed. The one real
world usage of this pattern is mentioned in this [wiki article](https://en.wikipedia.org/wiki/Double-ended_queue#Applications).

## Time Complexity

| Case      | Complexity  |
| --------- |:-----------:|
| All       | `O(1)`      |
