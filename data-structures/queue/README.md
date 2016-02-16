# Queue

A queue is a container which allows items to be pushed in and removed on a
first in first out principle (FIFO). In practice, we push new items onto
the back of the queue and pop items off the front of the queue.

In this implementation, I'm using a circular doubly linked list as the
container.

## Use Cases

In the real world, we could use a queue to simulate things such as the line of
cars at a traffic light, a printer buffer, packets waiting at a router. In
general terms, they can be used anywhere where you want to ensure that items
are executed in order they were created.

## Time Complexity

| Case      | Complexity  |
| --------- |:-----------:|
| All       | `O(1)`      |
