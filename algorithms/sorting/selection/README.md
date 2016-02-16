# Selection Sort

This is an example of Selection Sort. Selection sort does an in place sort
on a set (slice). It's easiest to explain how it works with an example:

If we have a set: `[3, 1, 6, 0, 5]`

We start at index 0 and see the number 3. We then go from index 1 through the
end of the set looking for the next smallest number. So:

```
1 < 6
1 > 0
0 < 5
```

This means the the number 0 at index 3 is the smallest number and we swap the
number 3 at index 0 with the number 0 at index 3 giving us:

```
[0, 1, 6, 3, 5]
```

We move onto index 1 where the value is 1 and check each number following it
like before:

```
1 < 6
1 < 3
1 < 5
```

This is the smallest number so our list is the same at:

```
[0, 1, 6, 3, 5]
```

We move onto index 2 where the value is 6 and we check each number following
it like before:

```
6 > 3
3 < 5
```

The number at index 3 is smaller than our 6 so we swap those values leaving us
with:

```
[0, 1, 3, 6, 5]
```

We move onto index 3 where the value is 6 and we check each number following
it like before:

```
6 > 5
```

The number at index 4 is smaller than our 6 so we swap those values leaving us
with:

```
[0, 1, 3, 5, 6]
```

On the last iteration, we see that this is the final item in the list so it
must be sorted and in fact, it is!

While this process is certainly tedious, it does give us a way to
programatticaly go through a set and sort it.

## Time Complexity

| Case      | Complexity  |
| --------- |:-----------:|
| Worst     | `O(n^2)`    |
| Best      | `O(n^2)`    |
| Average   | `O(n^2)`    |
