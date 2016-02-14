# Quick Sort

Quick sort is like merge sort in that it's a divide and conquer recursive
algorithm. It differs however by where the work is being done. In merge
sort, we do the bulk of the work in the combine step, but with quick sort
the work is done in the divide step. In general, it will outperform merge
sort and it is done in place which is even better. As an example:

If we have a set: `[3, 1, 6, 0, 5]`

Using quick sort, we find the **pivot** or point which we will split the set
into two sets. In practice, we select the last item in the set. This is the
**partition** phase. Based on this last value, 5 in our case, we visit each
number in the set checking if the value is higher or lower than 5. If it's
lower, then we will swap the current value with the minimum index. If it's
higher, we would swap the current value with the maximum index. In practice
here, this would look like:

```
5 is higher than 3, the set stays as [3, 1, 6, 0, 5]
5 is higher than 1, the set is [1, 3, 6, 0, 5]
5 is lower than 6, the set is [1, 3, 6, 0, 5]
5 is higher than 0, the set is [0, 1, 3, 6, 5]
There are no more values to compare so we move the pivot value to its place,
the set is [0, 1, 3, [5], 6].
```

As we can see, this is sorted with the pivot value being the split between
higher and lower items. This set is altered and we return the pivot index
to divide it again. We get:

`[0, 1, 3]` and `[6]`

The right side is sorted already since it's only one term. The left side,
when split, gives us 3 as the pivot value and when we check the items we
see that it's the highest so we are sorted there too. When we return this
and bubble up, we are sorted on both sides and combine the terms of:

```
[0, 1, 3, 5, 6]
```

This is efficient in a lot of ways.

## Time Complexity

| Case      | Complexity  |
| --------- |:-----------:|
| Worst     | `O(n^2)`    |
| Best      | `O(n lg n)` |
| Average   | `O(n lg n)` |
