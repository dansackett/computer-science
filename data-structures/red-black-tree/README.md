# Red-Black Tree

A red-black tree (RBT) is a self-balancing binary search tree allowing
constant time promises in all cases. Balance is achieved by using an extra bit
to label each node as either `red` or `black`. These colors are then used with
a set of rules dictating the balancing operations needed to be performed on
insert and delete. On each insert or delete, balancing is done to bring the
tree into an approximately balanced state.

While the balancing is not perfect, it is very good and remains as an
efficient way to handle binary search trees.

All nodes in a red-black tree have leaves (children) which are either null
pointers or ghost nodes. In most implmentations it's easiest to have these set
to an actual value rather than nil. In my implementation, I'm using a single
sentinel node in order to conserve memory.

These trees follow the normal BST properties with lesser numbers to the left
and greater numbers down the right. What sets it apart are the red-black
properties which are:

- A node is either red or black
- The root node is black (in a balanced state)
- All leaves are black (nil or ghost nodes)
- If a node is red, both of its children are black
- Every path from a single node to it's descendent leaves must contain the
  same number of black nodes.

Together these rules force the tree to have any path from the root to its
furthest leaf is only two times as long as the path from the root to the
nearest leaf.

## Terminology

When talking about a red-black tree or really any binary search tree, the
following terms will likely come up and should be understood:

- **Parent:** A parent node is the direct ancestor of the current node.
- **Grandparent:** A grandparent node is the direct ancestor of the current
  node's parent node.
- **Uncle:** An uncle node is the grandparent's other child.
- **Sibling:** A sibling node is the current node's parent's other child.

## Inserting New Nodes

Inserting into a red-black tree is the same as inserting into a BST. The only
exception is the callback behavior once the insert is finished. We must
rebalance the tree and we do so using the newly inserted node. We can base our
balancing actions on the following cases:

### Case 1: Current Node is Tree Root

If we're looking at the root of the tree, we simply need to repaint it black
to satisfy our second rule (root node is black). This doesn't affect the other
rules since it adds one black node to every route and therefore won't change
anything.

### Case 2: Current Node's Parent is Black

If the parent of our current node is black then inserting the new node didn't
affect anything since newly inserted nodes are red by default. This is valid
since it doesn't add an extra black node to any paths.

### Case 3: Parent and Uncle Nodes are Red

If the parent and uncle of our current node are both red then we can simply
repaint them black and subsequently paint the grandparent red to restore
balance for the subtree. This works because it means the grandparent has two
black children and our newly inserted red node doesn't affect any other black
paths in the tree.

After this, we need to ensure that we move up the tree to rebalance from the
grandparent node so we jump back to case 1 with the grandparent node as the
current node.

For an example of this:

![Red Black Tree - Case 3](https://github.com/dansackett/computer-science/blob/master/data-structures/red-black-tree/images/Red-black_tree_insert_case_3.svg)

### Case 4: Parent is Red and Uncle is Black

#### Subcase 1

In this subcase we will see one of two situations:

- Current node is its parent's left child and its parent is its grandparent's
  right child.
- Current node is its parent's right child and its parent is its grandparent's
  left child.

In both cases, the pattern to look for is the chain having alternating sides.
Our ideal situation is to have the current node, parent, and grandparent all
down the left or right side. To get there, we can do a left or right rotation
depending on the layout. The easy rule of thumb is the rotate around the
parent to the grandparent's leaning side. So for example:

Current node is on paren's left and parent is on grandparent's right. We do a
right rotation.

Once this rotation is done, we move into subcase 2 no matter what the result is.

![Red Black Tree - Case 4 - Subcase 1](https://github.com/dansackett/computer-science/blob/master/data-structures/red-black-tree/images/Red-black_tree_insert_case_4.svg)

#### Subcase 2

The final case to cover is when they are all leaning down the same side. For
instance, the current node is the parent's left child and the parent is the
grandparent's left child. We can recolor the grandparent as red, the parent
as black, and do a right rotation. If it's down the right side, we recolor the
same but do a left rotation.

![Red Black Tree - Case 4 - Subcase 2](https://github.com/dansackett/computer-science/blob/master/data-structures/red-black-tree/images/Red-black_tree_insert_case_5.svg)

## Deleting Nodes

Deleting nodes in a red-black tree is similiar to deleting in a standard BST
but there are a few differences. Like insertion, we can evaluate our deletion
algorithm based on a number of identifiable cases.

We first need to find the node in the tree to delete and we can do so by
recursing through the tree in order. If it's not in the tree, our job is done.
When we do find it, we label it as the current node as we did in the insertion
sequence.

We can now check how many children it has. If it has both a left and right
child, we must handle the case much like a binary search tree delete on two
children. We first find the smallest value in the right subtree and set the
current node's value to this value. Our new current node becomes the node we
replaced it with.

Since the smallest child would not have two children (it can't or else
something would be smaller) then we use this node as our new current node.

At this point, our current node should have one or no child nodes. If it has
one child node, we replace the current node with the child node to unlink our
node and effecively delete it from the tree. If that node was black and the
child was red then all we have to do to keep balanced is change the child's
color to black. If the node was black and the child was black though, we have
to begin balancing through cases (because removing a black node affects the
red-black rules for black paths).

### Case 1: Current Node has no parent

This would mean that we're at the root of the tree and that there is nothing
to rebalance at the root. Remember, the deleted node has already been deleted
at this point.

### Case 2: Sibling Node is Red

If the current node's sibling is red, we can get away with switch the parent
and sibling's colors and rotate around the parent. If the current node is the
left child, we do a left rotation. We do a right rotation in the opposite
case.

If the sibling is black, we move into case 3. The following shows the case of
a red sibling:

![Red Black Tree - Case 2](https://github.com/dansackett/computer-science/blob/master/data-structures/red-black-tree/images/Red-black_tree_delete_case_2.svg)

### Case 3: Parent Node is Black and Sibling Tree is Black

In the case that the parent is black and the sibling node and its two children
are also black then all we have to do is repaint the sibling node as red. We
then send the parent node back to case 1.

This is because when we deleted our node we altered the black path making the
sibling paths have an extra black node. If we repaint the sibling node to red
then we remove that extra black node and the paths are again correct.

If this case doesn't match we move into case 4. The following shows the above
case:

![Red Black Tree - Case 3](https://github.com/dansackett/computer-science/blob/master/data-structures/red-black-tree/images/Red-black_tree_delete_case_3.svg)

### Case 4: Parent Node is Red and Sibling Tree is Black

This case is similar to case 3 but instead of a black parent we have a red
parent. In this case, we can push the red color to the sibling node and set
the parent node to black to be back where we would be after case 3.

If the conditions don't match, you need to check case 5. The following shows
the above case:

![Red Black Tree - Case 4](https://github.com/dansackett/computer-science/blob/master/data-structures/red-black-tree/images/Red-black_tree_delete_case_4.svg)

### Case 5: Sibling Node is Red, Sibling's Right Child is Black, Sibling's Left Child is Red and Current Node is Parent's Left Child

In this case we want to get all items down the same side. This is similar to
the insertion case 4 and serves as a preparation step for case 6.
Nevertheless, the operations that need to be done are: Recolor the sibling
red, recolor the sibling's left child black, and rotate right around the
sibling.

Notice that for the right side we can flip the color cases above to do the
opposite balancing operations.

We can now move into the final case with this result. The following shows the
above case:

![Red Black Tree - Case 5](https://github.com/dansackett/computer-science/blob/master/data-structures/red-black-tree/images/Red-black_tree_delete_case_5.svg)

### Case 6: Sibling Node is Black, Sibling's Right Child is Red, Current Node is Parent's Left Child

In this final case, we will repaint the sibling with the parent's color and
repaint the parent as black. Since there is now an extra black ancestor we can
do a left rotation to fix it.

Notice that for the right side we can flip the color cases above to do the
opposite balancing operations.

The following shows the above case:

![Red Black Tree - Case 6](https://github.com/dansackett/computer-science/blob/master/data-structures/red-black-tree/images/Red-black_tree_delete_case_6.svg)

## Use Cases

Due to their efficiency, red-black trees can be used in a lot of applications
to help provide a constant time search, insert, and delete data structure.
Many data structures can use a balanced binary tree such as this as their
backbone. For instance, a heap or priority queue could use a balanced binary
search tree to provide fast lookups up of maximum and minimum items. In
general, places where searching is required a well-balanced BST is a good
place to look first.

## Time Complexity

| Case      | Average       | Worst Case    |
| --------- |:-------------:|:-------------:|
| Space     | `O(n)`        | `O(n)`        |
| Search    | `O(log n)`    | `O(log n)`    |
| Insert    | `O(log n)`    | `O(log n)`    |
| Delete    | `O(log n)`    | `O(log n)`    |
