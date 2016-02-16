# Linked List

A linked list is a collection of items, each called a node, which has a
reference pointer to the next node. Together, this group of nodes represents a
sequence. In the simplest form, a node stores both a value and a pointer to
the next node.

Linked lists are one of the simplest data structures and serves as a good base
to things like stacks, queues, and associative arrays. What makes a linked
list favorable is the ability to insert or delete new nodes without
reorganizing every node in the list.

The two types of linked lists are:

- Singly linked lists
- Doubly linked lists

## Singly Linked List

A singly linked list is the most basic form and is described such as:

```
[ 12 ] --> [ 66 ] --> [ 45 ] --> [ X ]
```

As can be seen, each node points to the next node where the last node (tail)
points to a nil reference.

## Doubly Linked List

A doubly linked list is an enhancement of the singly linked list in which it
holds both a reference to the next node and the previous node like so:

```
[ X ] <--> [ 12 ] <--> [ 66 ] <--> [ 45 ] <--> [ X ]
```

This gives us an even better way to operate and traverse the list while using
slightly more memory storing an additional pointer value.

## Circular Linked List

In the provided code, I have created a circular linked list which uses a
doubly linked list as the base but instead of pointing to nil at the head and
tail, they point to a **sentinel value** which is always locked to the 0
index. This allows us to work more efficiently with the list itself.

## Time Complexity

In most cases, this implementation gives us an O(1) time span. There is no
find() method currently so the only expensive method available is the
ToSlice() method which is O(n) as it must visit each node.

| Case      | Complexity  |
| --------- |:-----------:|
| Insert    | `O(1)`      |
| Delete    | `O(1)`      |
| Find      | `O(n)`      |
| Iterate   | `O(n)`      |
