# Merge Sort

Merge sort is a recursive sorting algorithm which utilizes the divide and
conquer concept. It works by taking a set of data, splitting it into two
halves, and merging those two halves in sorted order. For example:

If we have a set: `[3, 1, 6, 0, 5]`

We split that into two halves `[3, 1, 6]` and `[0, 5]`. We then split those
halves into halves again: `[3, 1]` and `[6]` and `[0]` and `[5]`. We split the last
half into two and then we begin the merge as we bubble back up:

```
                                                  --> [3]
                                    --> [3, 1] --|
                                   |              --> [1]
                   --> [3, 1, 6] --|
                  |                 --> [6]
[3, 1, 6, 0, 5] --|
                  |                 --> [0]
                   --> [0, 5   ] --|
                                    --> [5]
```

We can start from the final split and compare 3 with 1 and see that `1 < 3` so
we get the new set `[1, 3]`. Since we don't have the value to compare to `[6]`
yet we can check out the other split and compare `0 < 5` so we get `[0, 5]` as
we already had. Our new graphic would look like this for the next round:

```
                                    --> [1, 3]
                   --> [3, 1, 6] --|
                  |                 --> [6]
[3, 1, 6, 0, 5] --|
                   --> [0, 5   ]
```

On this round, we do a sort between our `[1, 3]` and `[6]`. We compare the
first indices of the sets and see that `1 < 6`. Since the left side won, we
remove the value at index 0 from the left side (1) and we then move onto
another comparison at index 0 between the new value on the left and the old
value on the right: `3 < 6`. We add 3 to the resulting set and add 6 onto the
end of the resulting set since it's the last number giving us `[1, 3, 6]`. Our
new graphic with what's left to merge looks like:

```
                   --> [1, 3, 6]
[3, 1, 6, 0, 5] --|
                   --> [0, 5   ]
```

This final merging step does as we did before comparing numbers at the 0
index. When one side wins, we remove the winning value from the set and
compare the next value with the losing sides value. So in short:

```
0 < 1
1 < 5
3 < 5
5 < 6
```

This gives us the final sorted set of `[0, 1, 3, 5, 6]`.

## Time Complexity

| Case      | Complexity  |
| --------- |:-----------:|
| Worst     | `O(n lg n)` |
| Best      | `O(n lg n)` |
| Average   | `O(n lg n)` |
