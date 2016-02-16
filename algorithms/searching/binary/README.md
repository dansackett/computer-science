# Binary Search

With Binary Search, the goal is to constantly work on smaller data sets in
order to limit the number of guesses one must make. Based on the size of
the array (slice) we're searching, we can quickly find an answer by
guessing the number in the middle of our maximum and minimum numbers. So
for instance, if we have:

```
[1, 2, 3, 4, 5, 6, 7]
```

And the number we're looking for is 5 then in a linear search it would
take us 5 guesses to get there. With binary search, it would only take us 3
guesses:

- We find that the middle index of the list is 3
- The value at index 3 is 4. Since our value we wish to find (5) is larger
  than 4, we discard the first half leaving the new set as `[5, 6, 7]`.
- We find the new middle index of the list is 5 (based on the original list)
- The value at index 5 is 6. Since our value we wish to find (5) is less than
  6, we discard the upper half of the set leaving the new set as `[5]`.
- The remaining middle index is at 4 and that value is 5, the number we are
  looking for.

## Time Complexity

| Case      | Complexity  |
| --------- |:-----------:|
| Find      | `O(lg n)`   |
