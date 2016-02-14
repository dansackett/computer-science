# Linked List

This is an example of a Circular Doubly-Linked List. This implementation
serves as a good and somewhat more efficient way to use a linked list in
programming. In a Singly-Linked List each node would have a reference to
only its next node but here, we use the Doubly-Linked List to keep
references to both previous and next nodes. This saves from iterating in a
number of cases and gives us much cleaner code with less checking.

This circular implementation makes use of a sentinel value locked in at the
0 index of our list. This sentinel points to the first and last nodes in
the list allowing us to:

- When iterating, know when we should be done iterating
- Avoid complicated spaghetti code setting the head and tail correctly

## Time Complexity

In most cases, this implementation gives us an O(1) time span. There is no
find() method currently so the only expensive method available is the
ToSlice() method which is O(n) as it must visit each node.
