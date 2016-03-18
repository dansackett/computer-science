# Set

A set is an abstract data type that stores unorder unique values. Sets differ
from other list-type containers because most operations of a set check for
existence rather than for actual values. Still, it can be used in place of an
array or another container if order doesn't matter and duplicates don't need
to be stored.

The typical two types of sets are **static sets** and **dynamic sets**. As you
can imagine the core difference is that statis sets only allow search and
traversal while a dynamic set can be added to and subtracted from.

## Implementation

Due to the set properties many sets are implemented using a hash table where
the keys are the values of the set. In my implementation here, I decided to
use a hash table for this to make it very easy to work with. Other structures
such as standard arrays and trees can be used if the tradeoffs are better for
your use case. For instance a balanced binary search tree will guarantee o(lg n)
complexity for most cases where a hash table may degrade to o(n) at worst
case.

A binary search tree may be overkill though because the ordering of a set
typically doesn't matter which is why a hash table can be a good candidate.

## Operations

Sets are based on math's finite sets so operations are very similar. Some of
the most common operations are:

- Union: Return a new set which is a combination of two sets.
- Intersection: Return a new set which contains the values that exist in two
  sets.
- Difference: Return a new set which contains the values that don't exist in
  another set. This will change based on the set asking for the difference.

Some other common operations are:

- IsSubset: Check if another set is the subset of your set.
- IsSuperset: Check if your set is the subset of another.

## Use Cases

A set can be used in places where membership is important rather than the real
contents of a container. For instance, if you simply want to see how two
containers compare then set operations can provide clarity. If you care about
unique data and don't care about order, a set can be a lighter way to store
items as well.

## Time Complexity

| Case      | Average       | Worst Case    |
| --------- |:-------------:|:-------------:|
| Space     | `O(n)`        | `O(n)`        |
| Search    | `O(1)`        | `O(n)`        |
| Insert    | `O(1)`        | `O(n)`        |
| Delete    | `O(1)`        | `O(n)`        |
