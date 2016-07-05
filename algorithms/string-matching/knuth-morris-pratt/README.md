# Knuth-Morris-Pratt String Searching Algorithm

The Knuth-Morris-Pratt (KMP) algorithm is a string searching algorithm which
attempts to find occurrences of a word in text based on information from
previous mismatches. It improves on a brute force approach by not checking
previously matched characters allowing for smart skipping through a set of
text.

The part that makes this algorithm better than some others is its' computation
of a **partial match table**. This table is used to determine where the next checks
begin in case of a match failure.

## The Partial Match Table

The partial match table, or failure function, is the piece that tells us where
to pick up if a mismatch occurs. Generating a partial match table was pretty
confusing to me until I read [this description](http://jakeboxer.com/blog/2009/12/13/the-knuth-morris-pratt-algorithm-in-my-own-words/).
I will attempt to summarize this further into my own explanation using the same
pattern he mentions (`ABABABCA`).

To start, we can look at the pattern like this for an idea about our end
result:

| Char  | A | B | A | B | A | B | C | A |
|:-----:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|
| Index | 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 |
| Value | ? | ? | ? | ? | ? | ? | ? | ? |

Looking at the table we can see that we lay out our pattern and match it to the
string index where it resides. The final row of values is the output of our
table and what we use in the search algorithm.

What's in this value row?

It's a computation based on proper prefixes and suffixes. The following should
explain what each is:

- **Proper Prefix:** The characters of a string where one or more is left off
  from the end. For example: `d, da` are both prefixes of `dan`.
- **Proper Suffix:** The characters of a string where one or more is left off
  from the beginning. For example: `n, an` are both suffixes of `dan`.

When looking for these things it's important to remember that we will iterate
through the string looking at different indices throughout the process. This
means that if we're looking at index 2 of our pattern in the table above then
we only mean the subpattern `ABA` and not the entire pattern.

With that in mind our goal for each index is to find the length of the longest
proper prefix that matches a proper suffix in the same subpattern. So when
looking at index 2 and the pattern `ABA` we can conclude that prefixes are `A,
AB` and suffixes are `A, BA`. Immediately we can see that `A` matches in each
and we can assign the value of 1 to this index:

| Char  | A | B | A | B | A | B | C | A |
|:-----:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|
| Index | 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 |
| Value | ? | ? | 1 | ? | ? | ? | ? | ? |

Looking at the table, you may be wondering about the values before this match.
Let's look at them quickly. Since the first character will always have the same
prefix and suffix (none) then we will usually assign it a value of -1. For the
second index we see there is no match so we can assign it a value of 0. This
will gives us a new table:

| Char  | A  | B | A | B | A | B | C | A |
|:-----:|:--:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|
| Index | 0  | 1 | 2 | 3 | 4 | 5 | 6 | 7 |
| Value | -1 | 0 | 1 | ? | ? | ? | ? | ? |

We can now move forward in the algorithm looking at index 3 now we can see that
the subpattern is `ABAB` with prefices of `A, AB, ABA` and suffixes of `B, AB,
BAB`. Since `AB` matches both then we can mark this index as a value of 2 since
the length of the match is 2:

| Char  | A  | B | A | B | A | B | C | A |
|:-----:|:--:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|
| Index | 0  | 1 | 2 | 3 | 4 | 5 | 6 | 7 |
| Value | -1 | 0 | 1 | 2 | ? | ? | ? | ? |

Now moving into index 4 we have the subpattern `ABABA` with prefixes `A, AB,
ABA, ABAB` and suffixes `A, BA, ABA, BABA`. Here we can see two matches which
makes this more interesting. In a case like this we choose the longer prefix
setting the value for index 4 as 3.

At this point the concept should make sense so we can finish the table and our
result is:

| Char  | A  | B | A | B | A | B | C | A |
|:-----:|:--:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|
| Index | 0  | 1 | 2 | 3 | 4 | 5 | 6 | 7 |
| Value | -1 | 0 | 1 | 2 | 3 | 4 | 0 | 1 |

This is great but how do we achieve this in code?

The essence will be using two pointers which tell us the current state of the
search. One pointer will move forward through the text incrementing on each
iteration. Out second pointer will always start at the front of the pattern. With
each iteration we check if the first pointer character matches the front of the
pattern. If it does then we have a match and need to see how long the match
runs for. In the case of a mismatch we need to move to the last known match
start so we can check the next substring.

While this is cofusing to explain in text, the code provided can clarify it a
bit.

## The Search Function

Once we have developed a partial match table we can use this to make our
algorithm more efficient. This efficiency comes from knowing if old comparisons
have already been done and effectively skipping those in favor of fresh
comparisons.

How does the table help?

When we find a partial match in the text we continue checking until we find a
mismatch or the pattern is found. In the case of a mismatch we use the partial
match table to decide where we look next. At this point we need to know the
length of the partial match that we found. We use this (minus 1 for indexing)
and that finds us the correct partial match entry. We can then move forward
`partial match length - table[partial match length - 1]`. Why this value? With
the partial match we need to ensure that we don't skip any other potential
matches. By looking up the value in our table we can see how many potential
matches we have up to this point. Subtracting that value from the actual
partial match should bring us to a safe place to continue the search.

With this I find that it can be harder to explain everything when the code
itself does a good job of giving understanding.


## Time Complexity

| Case      | Complexity  |
| --------- |:-----------:|
| Worst     | `O(2n)`     |
| Best      | `O(n)`      |
| Average   | `O(n + m)`  |
