# Shell Sort

Shell sort is an improvement on insertion sort which uses the concept of
comparing values that are far apart and reducing that gap until all values are
sorted. It is this gap sequence that defines the shell sort algorithm as the
gap determines the running time of the algorith. The challenging part about
gaps is that they are largely dependent on the input size.

To get an idea of how the process works we can think of these steps:

- Determine a reasonable gap
- Compare the element at an index to another element at index + gap
- If the elements are out of order, swap them over the gap
- Continue through the array swapping over the gap until you cannot
- Cut the gap in half and compare the values against this new gap
- The array is sorted after a gap of 1

The reasoning behind this is if you sort elements at a larger gap then the
array will become more sorted faster and later passes with smaller gap sizes
will be faster. This is the main difference from insertion sort since it
doesn't do as many swaps.

An example can help in describing this algorithm:

Let's assume we have the array: `[3, 7, 8, 5, 4, 1, 10]`

We can find a suitable gap by seeing the length of the array is 7. We can use
that knowlege to see `(7 / 2) + 1`. This will give us a gap which effectively
halves the data on first pass. This gap is `4`.

We can now compare elements at the gap positions:

**Pass 1**

[`3`, 7, 8, 5, `4`, 1, 10]  3 < 4 so no swap
[3, `7`, 8, 5, 4, `1`, 10]  7 > 1 so swap
[3, 1, `8`, 5, 4, 7, `10`]  8 < 10 so no swap

On second pass we determine the new gap which will be half of the previous gap
so we can use `2`.

**Pass 2**

[`3`, 1, `8`, 5, 4, 7, 10]  3 < 8 so no swap
[3, `1`, 8, `5`, 4, 7, 10]  1 < 5 so no swap
[3, 1, `8`, 5, `4`, 7, 10]  8 > 4 so swap
[3, 1, 4, `5`, 8, `7`, 10]  5 < 7 so no swap
[3, 1, 4, 5, `8`, 7, `10`]  8 < 10 so no swap

On the third pass, the new gap is `1` so we compare every value like a bubble
sort would.

**Pass 3**

[`3`, `1`, 4, 5, 8, 7, 10]  3 > 1 so swap
[1, `3`, `4`, 5, 8, 7, 10]  3 < 4 so no swap
[1, 3, `4`, `5`, 8, 7, 10]  4 < 5 so no swap
[1, 3, 4, `5`, `8`, 7, 10]  5 < 8 so no swap
[1, 3, 4, 5, `8`, `7`, 10]  8 > 7 so swap
[1, 3, 4, 5, 7, `8`, `10`]  8 < 10 so no swap

As we can see, the array is mostly sorted by the time we finish the second
pass and this is the reason for the shell sort over the insertion sort.

## Time Complexity

| Case      | Complexity  |
| --------- |:-----------:|
| Worst     | `O(n^2)`    |
| Best      | `O(n lg2 n)`|
| Average   | `Depends`   |
