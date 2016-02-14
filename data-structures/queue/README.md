# Queue

A queue is a container which allows items to be pushed in and removed on a
first in first out principle (FIFO). In practice, we push new items onto
the back of the queue and pop items off the front of the queue.

In this implementation, I'm using a circular doubly linked list as the
container.

## Time Complexity

Operations in a queue are O(1) as we simply push an item in and pop and
item off without any iteration or checking. This is especially true in this
implementation using a circular doubly linked list.
