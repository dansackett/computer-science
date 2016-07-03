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

## Tips and Tricks

**Determine if a number is odd or even**

If we want to determine if a number is odd or even we can simply check the
output of `x & 1`. The reasoning here is that odd numbers will always have the
least significant bit set. When we do an `&` with 1 then it should clear all
the bits if the number is even. If it's odd then it will clear all the bits
except for the least significant bit. Below is an example of this:

```
X = 43 = 00101011

43 & 1 = 00101011 & 00000001

  00101011
  00000001
& --------
  00000001
```

In the above example we can see that the result will be `00000001` since x has
the least significant bit already set.

**Test if the nth bit is set**

Building off the last hack we already know how to check the least significant
bit. If we want to check the nth bit then all we have to do is use a left shift
with it. Our manipulation is `x & (1 << n)`. If we think about this it makes
sense. Starting with 1 (00001) we shift the first bit over n times. When we do
the `&` operation then it should clear all bits except for the nth bit if it's
set.

Let's see an example considering the number 109 and if the 3rd bit is set:

```
109 = 1101101

1 << 3 = 0001000

x & (1 << 3) = 109 & 8 = 1101101 & 0001000

  1101101
  0001000
& -------
  0001000
```

In the above example we can see that the result is a non-zero number meaning
the third bit was set.

**Setting the nth bit**

On the flipside of the last hack, what if we want to ensure we set the nth bit?
We can do this in a similiar fashion by doing `x | (1 << n)`. If we think about
this it's almost identical to before but instead of flipping all bits with `&`
we're using `|` which will keep all bits and for the nth bit it will ensure
that it is flipped if it wasn't set. We can see another example using 109 as
the number and setting the 4th bit:

```
109 = 1101101

1 << 4 = 0010000

x | (1 << 4) = 109 | 16 = 1101101 | 0010000

  1101101
  0010000
| -------
  1111101
```

In the above example we can see that our 4th bit was flipped due to the `|`
operator essentially adding 16 to the number.

**Unset the nth bit**

As the compliment to the last hack we can also unset the last nth bit with
another similiar operation of `x & ~(1 << n)`. Since negating the left shift
will turn all bits but the nth bit to 1 and the nth bit to 0 then when we do
the `&` operation we will see the bits from x still set but if a bit exists in
x at position n then it will be flipped. Let's look at our example with the
number 109 again and unsetting the 3rd bit:

```
109 = 1101101

1 << 3 = 0001000
~(1 << 3) = 1110111

x & ~(1 << 3) = 109 & 119 = 1101101 & 1110111

  1101101
  1110111
& -------
  1100101
```

In the example above we can see that we successfully unset the nth bit turning
109 into 101.

**Toggle the nth bit**

If we want to switch the nth bit from whatever it is we can use the same
formula as the above example except using an xor `x ^ (1 << n)`. Remember with
xor that if both bits compared equal 1 then the result will be 0. If one is 1
and the other 0 then the result will be 1. With that in mind the n left shifts
will give us a value where xor will essentially do the toggling. Let's look at
an example using 109 and toggle the 3rd bit:

```
109 = 1101101

1 << 3 = 0001000

x ^ 1 << 3 = 109 ^ 8 = 1101101 ^ 0001000

  1101101
  0001000
^ -------
  1100101
```

In the above example our 3rd bit was flipped to 0. If we were to xor the result
with 1 << 3 again then it would flip that bit back proving the toggle.

**Turn off the rightmost 1 bit**

In a pattern the rightmost 1 bit is the smallest bit set. If we want to make
sure it is turned off then we can check `x & (x - 1)`. We looked at this
property in the first algorithm example but another example never hurts:

```
109 = 1101101

x - 1 = 108 = 1101100

x & (x - 1) = 109 & 108 = 1101101 & 1101100

  1101101
  1101100
& -------
  1101100
```

In the above example we see the result is the same pattern as 109 except the
rightmost 1 bit has been flipped. Coincidentily this equals the same as x - 1
but that's not always the case.

**Isolate the rightmost 1 bit**

Working with the rightmost 1 bit again we can also leave just that bit by doing
`x & -x`. This is a negative number which isn't the same as negating the
pattern. Before we see an example let's see how to make a number negative. You
may hear this called the "twos compliment" which roughly means `-x = ~x + 1`.
When dealing with twos compliment we will consider a leading 1 as a negative
number and leading 0 as positive. So for example we can see -9:

```
9  = 01001
-9 = ~9 + 1 = 10110 + 00001 = 10111
```

As we can see our representation of -9 negates x and adds 1. Now let's see how
this works for isolating the rightmost 1 bit. We will use a new number of 188:

```
188  = 10111100
-188 = 01000011 + 1 = 01000100

188 & -188 = 10111100 & 01000100

  10111100
  01000100
& --------
  00000100
```

As we can see in the example our end result has only the rightmost 1 bit set.
We now know this works as expected.

**Set all bits to the right of the rightmost 1 bit**

With this operation we find the rightmost 1 but in the pattern and set all bits
to the right of that equal to 1 as well. We do this by running `x | (x - 1)`.
For an example we can look at the number 188 again:

```
188  = 10111100
x - 1 = 187 = 10111011

188 | 187 = 10111100 | 10111011

  10111100
  10111011
| --------
  10111111
```

We know this is true because when we find x - 1 it flips all bits of x from the
rightmost 1 bit to the least significant bit. Because of this we know that
doing an `|` with x and x-1 will ensure that all bits right of the rightmost 1
but will be valid and set to 1 still. The example above shows this as well.
