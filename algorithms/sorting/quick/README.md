# Quick Sort

Quick sort is like merge sort in that it's a divide and conquer recursive
algorithm. It differs however by where the work is being done. In merge
sort, we do the bulk of the work in the combine step, but with quick sort
the work is done in the divide step. In general, it will outperform merge
sort and it is done in place which is even better. As an example:

If we have a set: `[3, 7, 8, 5, 4]`

Using quick sort, we find the **pivot** or point which we will split the set
into two parts. In practice, we select the last item in the set. This is the
**partition** phase. In our set, our first pivot is `4`. We then take this
pivot and compare it with each other number in the set:

```
4 > 3, set is: [3, 7, 8, 5, [4]]
4 < 7, set is: [3, 8, 5, [4], 7]
4 < 8, set is: [3, 5, [4], 8, 7]
4 < 5, set is: [3, [4], 5, 8, 7]
```

Now that the data is partitioned as `[3]` and `[5, 8, 7]` we can break them
apart and do a quick sort on each smaller set like we did before. In this
case, we see that `[3]` is alone so this is already sorted. Our other set is
not though. On that side, the pivot value becomes `7`.

```
7 > 5, set is: [5, 8, [7]]
7 < 8, set is: [5, [7], 8]
```

This gives us single values on each side as well and now we have the following
values that need to bubble up:

```
[3] (Initial left single number)
[4] (Initial pivot value)
[5] (Smaller of right subset)
[7] (Second pivot value)
[8] (Larger of right subset)
```

Bubbling up, we get the sorted set of `[3, 4, 5, 7, 8]`.

## Time Complexity

| Case      | Complexity  |
| --------- |:-----------:|
| Worst     | `O(n^2)`    |
| Best      | `O(n lg n)` |
| Average   | `O(n lg n)` |
