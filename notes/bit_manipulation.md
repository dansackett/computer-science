# Bit Manipulation

Bit manipulation has always been a bit confusing to me but I'm hoping to
outline some of the ideas here to keep them safe when I need a reference. In
most cases bit manipulation can help save memory and speed providing a low
level way to modify data. In some cases it can also simplify code quite a bit.

To start, let's look at a table to understand bits and bytes:

| From            | To              |
|:---------------:|:---------------:|
| 1 Bit           | 0 or 1 (Binary) |
| 8 Bits          | 1 Byte          |
| 1024 Bytes      | 1 Kilobyte (k)  |
| 1024 Kilobytes  | 1 Megabyte (mb) |
| 1024 Megabytes  | 1 Gigabyte (gb) |
| 1024 Gigabytes  | 1 Terabyte (tb) |
| 1024 Terabytes  | 1 Petabyte (pb) |

For bit manipulation we are working on the lowest possible level (bits) and a
binary 1 or 0 value. Since programming ultimately boils down to simple bits
then manipulating those values is incredibly close to the system itself which
explains the performance gains created by working at this level.

## Understanding Binary

While most people have seen this, it's good to highlight how a base 10 number
can be converted to a binary representation (base 2) and how to read that
binary number.

A binary number is a set of 1's and 0's in a pattern. The pattern is as
follows:

| 2^n | 2^4 | 2^3 | 2^2 | 2^1 | 2^0 |
|:---:|:---:|:---:|:---:|:---:|:---:|

I find that an example is best to explain how to use this pattern so we'll look
at the binary number 1101 as an example. To find out what this is in base 10 we
can start by separating the sequence of bits and matching them against the
above pattern like so:

| 2^3 | 2^2 | 2^1 | 2^0 |
|:---:|:---:|:---:|:---:|
| 1   | 1   | 0   | 1   |

With this table we can now build an equation where the bits are multiplied by
the column header:

```
(1 * 2^3) + (1 * 2^2) + (0 * 2^1) + (1 * 2^0)
```

If we solve this equation we will see that `1101` equals `13`. As you can see,
the pattern is the value and the bit tells us if that section of the pattern
should be applied to the total.

How do we go from decimal to binary though?

This isn't too bad actually. If we want to see what the number `9` is in binary
then we would start by thinking of the largest exponent of 2 that will not go
over our target. For this we can see that `2^3 = 8` is a good starting target.
We can now progressively start a table like so:

| 2^3 | 2^2 | 2^1 | 2^0 |
|:---:|:---:|:---:|:---:|
| 1   | ?   | ?   | ?   |

Moving down we see that adding the next slot will bring us over 9 so in there
we place a 0. We can also see that the next slot brings us over our target by
one meaning it also gets a 0. We then reach the final slot and it will put us
at 9 exactly meaning we mark it with a 1 leaving a final table of:

| 2^3 | 2^2 | 2^1 | 2^0 |
|:---:|:---:|:---:|:---:|
| 1   | 0   | 0   | 1   |

As we can see the number `9` in decimal is equivilent to `1001` in binary. It's
important to understand how to read and convert these numbers before getting
into any bit manipulation.

## Logical Operators

Bit manipulation hinges on the core concepts of logical idioms and
understanding a basic truth table and how logical operators work will be the
foundation of all manipulations. Below is a truth table for reference:

| A    | B    | !A    | A && B   | A &#124;&#124; B | A ^ B |
|:----:|:----:|:-----:|:--------:|:-----------------|:-----:|
| 0    | 0    |	1     |	0        | 0                |	0     |
| 0    | 1    |	1     |	0        | 1                |	1     |
| 1    | 0    |	0     |	0        | 1                |	1     |
| 1    | 1    |	0     |	1        | 1                |	0     |

The patterns in a truth table can be compared to bit operators almost directly
as can be seen in the table below:

| A    | B    | ~A       | A & B   | A &#124; B | A ^ B   |
|:----:|:----:|:--------:|:-------:|:----------:|:-------:|
| 1010 | 1100 |	0101     | 1000    | 1110       | 0110    |

Hopefully you can see that we work on each bit and can reference the truth
table to see the manipulation needed. For example look at `A & B`. Since A is
1010 and B is 1100 we can start at the front. `1 && 1 = 1` so our first value
is 1. Moving to the second values we look at `0 && 1 = 0` so our second value
is 0. Moving through the last two we get the final result as 1000. It's very
helpful to understand the truth table when working with bits.

## Bit (bitwise) Operators

In the last table I showed the four main bit operators but a short explanation
will probably help get the understanding across.

### NOT (~) Operator

The "not" operator is a unary operator (meaning it requires only itself) which
negates each bit to the opposite of itself. In essence it derives the
compliment of a bit. For example:

```
~1001 = 0110
```

### AND (&) Operator

The "and"operator is a binary operator which requires two equal-length bit
patterns. If both bits in the patterns are 1 then it will be a 1 but otherwise
it will be a 0. For example:

```
1001 & 1100 = 1000
```

### OR (|) Operator

The "or" operator is a binary operator which requires two equal-length bit
patterns. If either of the bits is a 1 then we get a 1. Only if neither bit is
a 1 do a we get a 0. For example:

```
1001 | 1100 = 1101
```

### XOR (^) Operator

The "xor" operator is a binary operator which requires two equal-length bit
patterns. If the bit in one pattern is a 1 and the other a 0 then we get a 1.
Otherwise if the bits match then we get a 0. For example:

```
1001 ^ 1100 = 0101
```

### Shift Operators

While the above four operators are the most common, shift operators are also
incredibly useful in bit manipulations. There are two types of shifts in bits:

- Left shift (<<)
- Right Shift (>>)

Shifts are unary operators which will shift the pattern in one direction.
Because of this shift we can consider a left shift to multiply the number by
2^n and a right shift to divide the number by 2^n. Let's look at some examples:

```
Left Shift
1 << 1 = 2 = 2^1
1 << 4 = 16 = 2^4

Right Shift
4 >> 1 = 2
16 >> 4 = 1
```

If we have actual binary numbers to shift then we can see how this is accounted
for:

```
Left Shift
0001 << 1 = 0010 = 2
0001 << 4 = 10000 = 16

Right Shift
0100 >> 1 = 0010 = 2
10000 >> 4 = 00001 = 1
```

## How They Help Algorithms

While it's good to see these things and understand them, seeing them in
practice is a different story. Let's look at some example problems to see how
they help us in performance and clarity.

### Check if a number is a power of 2

To solve this we could create a loop which divides the number by two until it
either is 1 or if it isn't then it is not a power of two. The following is
sample code for that solution:

```
func isPowerOfTwo(x int) {
    for x % 2 == 0 {
        x = x / 2
    }
    return x == 1
}
```

In this approach we are looking at a complexity of O(log n) since we constantly
cut the vector in half. Can we do better?

We can actually with bit manipulation but will require us to get a little
clever. One of the properties of binary numbers is that for a number x, x-1 can
be found by finding the rightmost 1 bit and fliping itself and all values to
the right of it. For example:

```
X = 6 = 110
X - 1 = 5 = 101

6 = 110
5 = 101
```

In the above example we find the rightmost 1 bit (the center 1) and flip itself
and all values to the right of it giving us 101 or 5. Let's see it with a
larger number:

```
X = 19
X - 1 = 18

19 = 10011
18 = 10010
```

Knowing this property we can also discover that `x & (x - 1)` is another
pattern where we see that the result will be almost identical to x itself. The
only difference is that the rightmost 1 bit in x will be flipped in the result.
For example:

```
X = 6 = 110
X - 1 = 5 = 101

110 & 101 = 100
```

One last thing that we need to recognize with binary numbers that are powers of
two are that they share one characteristic: only one 1 bit exists in the
pattern:

```
2  = 00010
4  = 00100
8  = 01000
16 = 10000
```

This final piece of the puzzle should give you a clue on how our algorithm will
work. Since `x & (x - 1)` will have the same pattern as x with the rightmost 1
bit flipped then any power of 2 will have no 1 bits present at all. That said
our algorithm will look like this:

```
func isPowerOfTwo(x int) {
    return (x && !(x & (x - 1)))
}
```

If the value is a power of 2 then negating the result will be truthy. We do a
check to ensure that the passed in integer is also truthy and if so then a
power of 2 will return true. If we negate the phrase and have a falsy value
though then we will return true.

With bit manipulation we took a logarithmic algorithm and made it a constant
time algorithm.
