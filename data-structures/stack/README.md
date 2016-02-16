# Stack

A stack is a container which allows items to be pushed in and removed on a
last in first out principle (LIFO). In practice, we push new items onto
the back of the stack and pop items off the back of the stack.

In this implementation, I'm using a circular doubly linked list as the
container.

## Use Cases

Stacks are used in a number of concepts. One of the first real world cases is
in writing reverse polish notation. Another good place to see a stack is in
browser history where if you go back to the last page, it would be the most
recent page pushed onto the stack. In general though, applications which could
recall the most recent value are good candidates for a stack.

## Time Complexity

| Case      | Complexity  |
| --------- |:-----------:|
| All       | `O(1)`      |
