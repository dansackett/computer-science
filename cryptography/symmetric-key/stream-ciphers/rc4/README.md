# RC4 Cryptographic Algorithm

The RC4 algorithm today is not a secure form of cryptography but it does
provide a good way to learn about some cryptographic ideas. In its simplest
description, it generates a pseudorandom stream of bytes which can be used to
encrypt plaintext into a ciphertext.

When starting with RC4 we first must initialize a state. The state has two
parts to it:

- A permutation of 256 possible bytes
- Two 8-bit index pointers

The permutation is initialized with a key typically between 40 and 256 bytes.
We use a **key-scheduling** algorithm we can randomize the bytes through 256
rounds of swaps. The result is a set of bytes which can be used in a
**pseudo-random generation algorithm**. The result from this is a stream of
bits which can be used to XOR against our plaintext to develop ciphertext.

## Key Scheduling Portion

Key scheduling helps us setup our permutation of 256 bytes pool based on a
given key. We first start by initializing our array:

```
var S [256]int
for i := 0; i < 256; i++ {
    s[i] = i
}
```

Once initialized then we can send the result through 256 rounds of iteration.
Here is where we randomize the bytes based on the key size. We do this like so:

```
j := 0
for i := 0; i < 256; i++ {
    j = (j + key[i % key_length] + S[i]) % 256
    S[i], S[j] = S[j], S[i]
}
```

We can see that for each index `i` in `S[0..255]` we determine an index `j` to
swap places with. This swap index is based on the previous iteration of itself
plus the byte from the key located at `i % key_length` plus our permutation
index at `i`. We then make sure it's a valid index in `S` by getting a modulo
of 256b (our bytes length).

Why like this?

While I don't know why exactly I can guess that it is a good way to randomize
each swap index. Our key is the basis of this randomness so initializing the
cipher with the same key again will gives us a predictable array of bytes
again.

## Pseudo-Random Generation Portion

Once we have our pool of randomized 256 bytes we can use it to develop a
keystream. The keystream is what we use to encrypt plaintext. In a typical
implementation (and mine) we will generate a byte for each byte of the
plaintext. We can do this like:

```
i, j = 0, 0
for v := range src {
    i = (i + 1) % 255
    j = (j + c.S[i]) % 255
    c.S[i], c.S[j] = c.S[j], c.S[i]
    output c.S[(c.S[i] + c.S[j]) % 255]
```

We first initialize our indices to 0 for a fresh start. We then find our first
index `i` by incrementing it by one (and a modulo by 255 to not overflow the
8-bit integer size). We then determine the index `j` by adding the byte located
at `S[i]` to j. We swap those values in the byte pool and then find the random
byte from the pool located at `(S[i] + S[j]) % 255`. This random byte is a
value of our keystream.

At this point we can XOR our src byte with this value of the keystream and get
the first resulting byte of our ciphertext.

We continue this until we have generated all of our keystream or XORed the
entire source string.

## More Information


