# Asymptotic Notation

When dealing with algorithms, it's important to understand the complexity and
running time of them in order to make decisions about efficiency. In most
cases, there are three different parts of Asymptotic Notation including:
big-Θ(Theta), big-O, and big-Ω(Omega).

## Big-Θ

With big-Θ (big theta) we can easily deduce a problem by reasoning that the
constant factors become negligible at a large point. So for instance, in the
equation `6n^2 + 100n + 300` we can see that `n^2` is going to be the highest
term in the sequence and at some point the other numbers will not affect the
equation in a significant way. So based on that, we can drop the `100n + 300`
and we're left with just `6n^2`. Even here, the exponent will make the
multiplier insignificant at some point so we can drop this as well and we're
left with just `n^2`.

We can then say that this equation (algorithm) is Θ(n^2).

When we say this, we mean that for a large bound of n, we can see that our
complexity will be between an upper bound of `k2 * n^2` and a lower bound of
`k1 * n^2`. This takes into account compiler time, the language, the compiler
being used, etc and gives us a fair estimation of what we can expect.

Some samples or common big-theta notation and their types are:

**Constant**

- `Θ(1)`

**Linear**

- `Θ(lg n)`
- `Θ(n)`
- `Θ(n lg n)`

**Polynomial**

- `Θ(n^2)`
- `Θ(n^2 lg n)`
- `Θ(n^3)`

**Exponential**

- `Θ(2^n)`

## Big-O

Where Big-Theta gives us a value for an algorithm both above and below, Big-O
notation only gives us an estimation based on the upper bound value. In the
computer science world, we use Big-O to describe the best case scenario at
times. For instance, with Binary Search we can see that saying it runs
`Θ(lg n)` is not always correct since it may be more efficient than that. For
instance, if we guess the value on the first iteration, then we are looking at
performance as `Θ(1)`. For this reason, we can use Big-O notation to note the
worse case scenario with probability that it can run better.

So we can say best that binary search is `O(lg n)`.

## Big-Ω

In contrast to Big-O, we can say that Big-Omega gives us the least amount of
time for an algorithm to run. This means we have just a lower bound and no
upper bound. With this, we can make assertions that binary search is in fact
`Ω(1)` meaning that binary search is at least constant.

These types of assumptions are not precise but still give us a good way to
define an algorithm.
