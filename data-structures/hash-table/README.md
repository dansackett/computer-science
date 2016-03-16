# Hash Table

A hash table is a data structure which maps keys to values. It uses a hash
function which finds an index based on the key and stores the value in a
"bucket" or "slot". In an ideal situation, each bucket will only contain one
value but this is not always the case. In the instance of a collision two keys
can reside in the same bucket which can ultimately degrade the hash table
performance to O(n) at worst case.

When building a hash table it's important to maintain a good balance between
the number of entries and number of buckets used. We can determine this "load
factor" by dividing the number of entries by the number of buckets. A larger
load factor points to a slower hash table.

## Hash Functions

Having a good hash function is incredibly important to a hash table. Hash
functions determine how to distribute values into the hash table and a poor
distribution will lead to fast degradation.

A common pattern in hashing will follow the formula `index = f(key, array_size)`.

Our function `f()` then could be:

```
hash = hashfunc(key)
index = hash % array_size
```

This concept allows the hash to be independent of the size of the underlying
array.

It's incredibly hard to have a perfect hash function without knowledge of the
input but in the case that all keys are known ahead of time it is possible to
build a "perfect hash function" which has no collisions.

## Collision Resolution

Collisions are almost guaranteed to happen in a hash table with an unknown set
of keys. Luckily there are a number of strategies that can be implemented in
order to mitigate collisions.

### Separate Chaining

Separate chaining is the process of storing matching keys in a list-type data
structure. In this case finding the bucket can be done in constant time and
finding the correct key is dependent on the structure used within. It should
be said that the structure you use shouldn't be one that expects to store a
large number of items. If your bucket sizes are that large then there is
something wrong with the hash function you're using and should evaluate there
first.

A linked list is a common structure for storing these values. They can be used
to traverse the bucket based on linked-list complexity times. While there is
some overhead storing a pointer to the next node they still are a valid option
for handling this chaining.

Balanced binary search trees can be another good option but the needs of the
system should be thought about first. While this will give the bucket an O(lg n)
lookup time it may take more memory for that guarantee.

### Open Addressing

Open addressing stores all entries in the bucket array itself and inserts new
entries based on a probing sequence. It then uses this same probing sequence
when searching for entries later. Some of the well known probing sequences
include:

- Linear Probing: Interval between probes is fixes (usually 1)
- Quadratic probing: Interval between probes is increased by adding the
  successive outputs of a quadratic polynomial to the starting value given
  by the original hash computation
- Double hashing: Interval between probes is computed by a second hash function

Unfortunately this approach may require dynamic resizing since a growing load
factor makes this almost unusable. On the other hand it decreases the time
needed to allocate each new entry.

## Dyanamic Resizing

Dynamic Resizing is the act of resizing the hash table itself as the load
factor grows. In a lot of applications, a loadfactor of 2/3 is a good point to
consider a resize so the hash table retains it's optimized state. Since
inserts are the only time a table becomes larger the resizing would be run at
this stage. Since changing the size of the underlying array will result in
hashes changing, resizing also includes rehashing or rebuilding the existing
hashes.

In many cases the entire table can be copied to a newly allocated array and
still perform well enough. In other cases such as real-time systems though
this penalty can be costly and an incremental approach needs to be taken. This
approach may follow a pattern:

- Allocate a new hash table and keep the old table
- Inserts will be placed in the new hash table
- Lookups will check both tables as they exist
- After an insert, a set number of entries will be copied to the new hash
  table
- Deallocate old hash table once all entries have been copied over

## Use Cases

Because most cases of the hash table have such fast lookups, hash tables are
used incredibly often. Some of the common use cases for a hash table may be an
associative array, database indexing, caches, and sets.

## Time Complexity

| Case      | Average       | Worst Case    |
| --------- |:-------------:|:-------------:|
| Space     | `O(n)`        | `O(n)`        |
| Search    | `O(1)`        | `O(n)`        |
| Insert    | `O(1)`        | `O(n)`        |
| Delete    | `O(1)`        | `O(n)`        |
