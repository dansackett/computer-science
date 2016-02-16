# Stack

A stack is a container which allows items to be pushed in and removed on a
last in first out principle (LIFO). In practice, we push new items onto
the back of the stack and pop items off the back of the stack.

In this implementation, I'm using a circular doubly linked list as the
container.

## Time Complexity

Operations in a stack are O(1) as we simply push an item in and pop and
item off without any iteration or checking. This is especially true in this
implementation using a circular doubly linked list.

| Case      | Complexity  |
| --------- |:-----------:|
| All       | `O(1)`      |
