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
binary 1 or 0 value. So as you can see, getting this low can be quite fast and
efficient.

With bits the common operations are:

- & (and)
- | (or)
- ~ (not)
- ^ (xor)

Before looking at how those operations work we can first look at a normal truth
table based on boolean values:

| `A`  | `B`  | `!A`  | `A && B` | `A || B` | `A ^ B` |
|:----:|:----:|:-----:|:--------:|:--------:|:-------:|
| 0    | 0    |	1     |	0        | 0        |	0       |
| 0    | 1    |	1     |	0        | 1        |	1       |
| 1    | 0    |	0     |	0        | 1        |	1       |
| 1    | 1    |	0     |	1        | 1        |	0       |

When working with bits we follow a lot of these same patterns as shown in the
following bit manipulation table:

| `A`  | `B`  | `~A`     | `A & B` | `A | B` | `A ^ B` |
|:----:|:----:|:--------:|:-------:|:-------:|:-------:|
| 1010 | 1100 |	11110101 | 1000    | 1110    | 0110    |

Hopefully you can see that we work on each bit and can reference the truth
table to see the manipulation needed. For example look at `A & B`. Since A is
1010 and B is 1100 we can start at the front. `1 && 1 = 1` so our first value
is 1. Moving to the second values we look at `0 && 1 = 0` so our second value
is 0. Moving through the last two we get the final result as 1000. It's very
helpful to understand the truth table when working with bits.

To be continued...
