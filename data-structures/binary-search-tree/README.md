# Binary Search Tree

A binary search tree (BST) is a data structure which allows fast lookups and
addition and removal of items. A BST's keys remain in sorted order so binary
search can be used when traversing a tree. Since we can skip different roots
of the tree, traversal is much faster than an unsorted array.

Every BST starts with a root node. Root nodes point to left and right subtrees
with the property that nodes with values less than the root are stored to the
left and nodes with values greater than the root are to the right. For
example:

```
     [ 5 ]
    /     \
   /       \
[ 3 ]     [ 10 ]
```

**Insertion**

When inserting nodes in a BST, we need to remember to always follow the
principle of "smaller values on the left, larger on the right". Using images,
let's start with an empty BST and insert a value of `5`:

```
[ 5 ]
```

This value is now our root node since the BST was previously empty. If we want
to add another value such as `10`, we would do so by starting at the root node.
We compare the value at the root node (`5`) to the new node value (`10`) and
see that our new node is greater so we add it to the right subtree:

```
[ 5 ]
     \
      \
     [ 10 ]
```

This represents the BST property and shows why searching for values is quick.
Since we know that values to the right are always greater than the current
node, we can continue to look down the right subtree to find our node.

If we want to add another node such as `3` then we would compare `3 < 5` and
see that we would add it to the left subtree like so:

```
     [ 5 ]
    /     \
   /       \
[ 3 ]     [ 10 ]
```

As a last look, let's add a few nodes and see how they would be arranged.
We're going to add `1`, `8`, and `20`:

```
          [     5    ]
         /            \
        /              \
     [ 3 ]           [ 10 ]
    /               /      \
   /               /        \
[ 1 ]           [ 8 ]     [ 20 ]
```

As we can see, if we wanted to find `20` it would only take 3 iterations where
a normal sorted array would take 6 iterations.

**Deletion**

When deleting, we have three cases to consider in order to keep the integrity
of the BST:

- Nodes with two children
- Nodes with on child
- Nodes with no children

Obviously, the most complex deletion involves two children so I'll start with
the easiest: no children. If we were to delete `1` from the final tree above
then our tree would look like:

```
          [     5    ]
         /            \
        /              \
     [ 3 ]           [ 10 ]
                    /      \
                   /        \
                [ 8 ]     [ 20 ]
```

Since deleting `1` does not affect any children, all we have to do is update
the parent node's (`3`) left subtree to nil.

Using the initial tree again, we can see what deleting `3` would simulate a
node having only one child. In this case, we check to see which subtree is the
child. In our case, it would be the left where `1` exists. So deleting `3`
would then need to link `1` to the root node `5` giving us:

```
          [     5    ]
         /            \
        /              \
     [ 1 ]           [ 10 ]
                    /      \
                   /        \
                [ 8 ]     [ 20 ]
```

Using the initial tree again, we can see that deleting `10` would simulate a
node having both children. With this, we first must find the smallest node
that exists in the right subtree. For us, this is `20`. We now can update node
`10` to be `20` while preserving the link to the root and the left subtree. We
then make sure we go down the right subtree to delete the exchanged value (`20`).
Our tree would now look like:

```
          [     5    ]
         /            \
        /              \
     [ 3 ]           [ 20 ]
    /               /
   /               /
[ 1 ]           [ 8 ]
```

As we can see, our BST property remains in tact this way.

**Traversing**

When traversing a BST, there are three main methods used:

- In Order Traversal: This is proper sorted order. It will recurse down the
  left subtree first, then the root of that tree, and then the right.
- Pre Order Traversal: This is a way to get the root nodes first. For our tree
  it would give us `5, 3, 1, 10, 8, 20`.
- Post Order Traversal: This is a way to get the left and right nodes first.
  For our tree it would give us `1, 3, 8, 20, 10, 5`.

## Use Cases

Binary Search Trees are good for a number of cases but at its worst case, it's
no better than a linked list (where all nodes are greater). If you plan on
incrementally adding and deleting values though, a BST is a great data
structure to keep that sequence sorted at all times.

Another good use case of a BST is the basis of a min or max heap. Since we can
easily recurse subtrees in logarithmic time then we can remove the min or max
value to satisfy the heap.

A deque is another data structure which, like the heap, makes a good candidate
for an underlying BST.

## Time Complexity

| Case      | Average       | Worst Case    |
| --------- |:-------------:|:-------------:|
| Space     | `O(n)`        | `O(n)`        |
| Search    | `O(log n)`    | `O(n)`        |
| Insert    | `O(log n)`    | `O(n)`        |
| Delete    | `O(log n)`    | `O(n)`        |
