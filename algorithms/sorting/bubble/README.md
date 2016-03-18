# Bubble Sort

Bubble sort is a brute-force sorting algorithm which repeatedly steps through
a list comparing adjacent elements and swapping them if in the wrong order. It
continues to do this pass through until no more swaps are needed. Bubble sort
in almost all cases is a poor choice for sorting in real world applications
but since the algorithm doesn't manipulate anything on a pass where the data
is already sorted it is good to verify that an array is sorted. This differs
from other sorting techniques since they actually do their full sorting
behavior since it doesn't know differently.

An example of operations is as follows:

**First Pass**

( `5` `1` 4 2 8 ) -> ( `1` `5` 4 2 8 ), Here, algorithm compares the first two elements, and swaps since 5 > 1.
( 1 `5` `4` 2 8 ) -> ( 1 `4` `5` 2 8 ), Swap since 5 > 4
( 1 4 `5` `2` 8 ) -> ( 1 4 `2` `5` 8 ), Swap since 5 > 2
( 1 4 2 `5` `8` ) -> ( 1 4 2 `5` `8` ), Now, since these elements are already in order (8 > 5), algorithm does not swap them.

**Second Pass**

( `1` `4` 2 5 8 ) -> ( `1` `4` 2 5 8 )
( 1 `4` `2` 5 8 ) -> ( 1 `2` `4` 5 8 ), Swap since 4 > 2
( 1 2 `4` `5` 8 ) -> ( 1 2 `4` `5` 8 )
( 1 2 4 `5` `8` ) -> ( 1 2 4 `5` `8` )

Now, the array is already sorted, but the algorithm does not know if it is completed. The algorithm needs one whole pass without any swap to know it is sorted.

**Third Pass**

( `1` `2` 4 5 8 ) -> ( `1` `2` 4 5 8 )
( 1 `2` `4` 5 8 ) -> ( 1 `2` `4` 5 8 )
( 1 2 `4` `5` 8 ) -> ( 1 2 `4` `5` 8 )
( 1 2 4 `5` `8` ) -> ( 1 2 4 `5` `8` )

## Time Complexity

| Case      | Complexity  |
| --------- |:-----------:|
| Worst     | `O(n^2)`    |
| Best      | `O(n)`      |
| Average   | `O(n^2)`    |
