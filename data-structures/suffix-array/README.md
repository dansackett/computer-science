# Suffix Array

A suffix array is basically just a sorted array of all the possible suffixes
of a given string. We can think of a suffix as a "substring" in this case. A
suffix array will contain indices to each suffix found in a string after they
have been sorted. In a lot of implementations, it's common to use a sentinel
character which is not in the alphabet of the string in question.

Let's look at an example to understand how this works using the word `banana`.

Using a sentinel value of `$`, we can see the string to index becomes `banana$`.
We index these items as follows:

| i     | 1 | 2 | 3 | 4 | 5 | 6 | 7 |
|:-----:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|
| S[i]  | b | a | n | a | n | a | $ |

With this, we can then start to break down the suffixes of this string such
that:

| Suffix    | i |
|:---------:|:-:|
| banana$   | 1 |
| anana$    | 2 |
| nana$     | 3 |
| ana$      | 4 |
| na$       | 5 |
| a$        | 6 |
| $         | 7 |

Finally, we can now sort these suffixes in ascending order to give us:

| Suffix    | i |
|:---------:|:-:|
| $         | 7 |
| a$        | 6 |
| ana$      | 4 |
| anana$    | 2 |
| banana$   | 1 |
| na$       | 5 |
| nana$     | 3 |

Our final suffix array with the corresponding indices to the suffixes (A) would be:

| i     | 1 | 2 | 3 | 4 | 5 | 6 | 7 |
|:-----:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|
| A[i]  | 7 | 6 | 4 | 2 | 1 | 5 | 3 |

## Suffix Array Construction

A proper suffix array can be constructed many ways. The way that I chose to
follow has a time complexity of `O(n log^2 n)`. The process is similiar to
above but continuously updates the order based on sets of two characters. An
example is best to understand this process:

| Sort Index    | Suffix Index  | Suffix            |
|:-------------:|:-------------:|:-----------------:|
| 0             | 10            | **i**             |
| 1             | 7             | **ip**pi          |
| 2             | 1             | **is**sissippi    |
| 2             | 4             | **is**sippi       |
| 3             | 0             | **mi**ssissippi   |
| 4             | 9             | **pi**            |
| 5             | 8             | **ppi**           |
| 6             | 3             | **si**ssippi      |
| 6             | 6             | **si**ppi         |
| 7             | 2             | **ss**issippi     |
| 7             | 5             | **ss**ippi        |

The table above shows the suffixes of the string `mississippi` sorted by the
first two characters in the suffix. For suffixes that have the same first two
characters, we assign the same sort index. With this, we can now create a
tuple to help us sort suffixes with the same index. We do this as such:

| Sort Index    | Suffix Index  | Suffix            | Suffix Index (After first 2 chars)    | Sort Index Tuple  |
|:-------------:|:-------------:|:-----------------:|:-------------------------------------:|:-----------------:|
| 0             | 10            | **i**             | -1                                    | (0, -1)           |
| 1             | 7             | **ip**pi          | 9                                     | (1, 4)            |
| 2             | 1             | **is**sissippi    | 3                                     | (2, 6)            |
| 2             | 4             | **is**sippi       | 6                                     | (2, 6)            |
| 3             | 0             | **mi**ssissippi   | 2                                     | (3, 7)            |
| 4             | 9             | **pi**            | -1                                    | (4, -1)           |
| 5             | 8             | **ppi**           | 10                                    | (5, 0)            |
| 6             | 3             | **si**ssippi      | 5                                     | (6, 7)            |
| 6             | 6             | **si**ppi         | 8                                     | (6, 5)            |
| 7             | 2             | **ss**issippi     | 4                                     | (7, 2)            |
| 7             | 5             | **ss**ippi        | 7                                     | (7, 1)            |

Based on the first sort, we look at the suffix that follows the first two
characters. As an example, let's look at the suffix `ippi`. The remaining
suffix after the first two characters is `pi`. We can see that this correlates
to the suffix at index 9 which has a sort index of 4. We create a tuple with
the initial suffixes sort index (`1`) and the remaining sort index (`4`)
giving us a value that we can use to sort things. As you can probably see, if
there is no remainder we assign a -1 as the remaing sort index.

When we do the sort on this tuple, our comparison function works something
like this:

```
Using the first two tuples (0, -1) and (1, 4):

if 0 == 1:
  return -1 < 4
else:
  return 0 < 1

For this, we would take the else route and since the result is true, they are
already in the correct order.

As an example of an order that would change, we can look at (6, 7) and (6, 5)
which correlate to the suffix indexes 3 and 6 respectively.

if 6 == 6:
  return 7 < 5
else:
  return 6 < 6

Since we are within the initial if statement and return false, they switch order.
```

So now when we sort the suffixes based on these results, we have all the
suffixes sorted by the first 4 characters. Our new result would look like:

| Sort Index    | Suffix Index  | Suffix            |
|:-------------:|:-------------:|:-----------------:|
| 0             | 10            | **i**             |
| 1             | 7             | **ippi**          |
| 2             | 1             | **issi**ssippi    |
| 2             | 4             | **issi**ppi       |
| 3             | 0             | **missi**ssippi   |
| 4             | 9             | **pi**            |
| 5             | 8             | **ppi**           |
| 6             | 6             | **sipp**i         |
| 7             | 3             | **siss**ippi      |
| 8             | 2             | **ssis**sippi     |
| 9             | 5             | **ssip**pi        |

On the next iteration, we move the "cursor" forward and can construct the same
set of tuples for the first 8 characters. This process continues until we
reach the end of the suffix length available.

In my algorithm, followed [this paper](http://web.stanford.edu/class/cs97si/suffix-array.pdf).

## Use Cases

In a lot of ways, this is used in contest-type problems dealing with finding
substrings. On a very practical level, it can be used to do full text
indexing, data compression algorithsm, and bioinformatics. In a more common
form, you will see these types of suffix arrays presented as suffix trees (a
form of trie) but the array version is a condensed way to store suffix
positions without using memory to save references to nodes.

Things to look for when deciding if a suffix array is right for the problem
are:

- Finding the longest common prefix of a string
- Finding if a substring exists in a string
- Searching through text

## Time Complexity

| Case      | Complexity        |
| --------- |:-----------------:|
| Construct | `O(n log^2 n)`    |
| Search    | `O(m log n)`      |
