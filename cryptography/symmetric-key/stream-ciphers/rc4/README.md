# RC4 Cryptographic Algorithm

The RC4 algorithm today is not a secure form of cryptography but it does
provide a good way to learn about some cryptographic ideas. In its simplest
description, it generates a pseudorandom stream of bytes which can be used to
encrypt plaintext into a ciphertext.

When starting with RC4 we first must initialize a state. The state has two
parts to it:

- A permutation of 256 possible bits
- Two 8-bit index pointers

The permutation is initialized with a key typically between 40 and 256 bits.
We use a **key-scheduling** algorithm to randomize the bits through 256
swaps. We can then use those shuffled bits in a **pseudo-random generation algorithm**.
The end result is a keystream that can encrypt each bit of the plaintext into a
ciphertext.

## Key Scheduling Portion

Key scheduling helps us setup our permutation of 256 bit pool based on a
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
swap places with. This swap index is based on the previous state
plus the bit from the key located at `i % key_length` plus our permutation
index at `i`. We then make sure it's a valid index in `S` by getting a modulo
of 256 (our bits length).

Why like this?

Using the key we can make a reproducible randomization of our bits. If we were
to initialize the cipher again with the same key then we would get the same
result making is predictable for reproducing a keystream. Each different key
will give us different permutations leading to a different keystream.

## Pseudo-Random Generation Portion

Once we have our pool of randomized bits we can use it to develop a
keystream. The keystream is what we use to encrypt plaintext. In a typical
implementation (and mine) we will generate a bit for each bit of the
plaintext. We can do this like:

```
i, j = 0, 0
for v := range src {
    i = (i + 1) % 255
    j = (j + S[i]) % 255
    S[i], S[j] = S[j], S[i]
    output S[(S[i] + S[j]) % 255]
```

We first initialize our indices to 0 for a fresh start. We then find our first
index `i` by incrementing it by one (and a modulo by 255 to not overflow the
8-bit integer size). We then determine the index `j` by adding the bit located
at `S[i]` to j. We swap those values in the bit pool and then find the random
bit from the pool located at `(S[i] + S[j]) % 255`. This random bit is a
value of our keystream.

At this point we can XOR our source bits with this value of the keystream and get
the first resulting byte of our ciphertext.

We continue this until we have generated all of our keystream or XORed the
entire source string.

## An Example

It may help to see an example of the process. In this example we'll use the
secret key of `hello` and a plaintext of `Attack at dawn`.

We first create our bit pool of `S[0..255]`. We then begin key scheduling by
starting our swap loop.

**First Iteration**

```
i = 0
j = 0

j = (j + key[i % key_length] + S[i]) % 255
j = (0 + key[0 % 5] + 0) % 255
j = (0 + "h" + 0) % 255
j = 104 % 255
j = 104
```

Above I showed how the equation is solved in each step. The one thing to note
here is that we had to convert `h`, our key value at index 0, to its ordinal
value (104). Once we have j we then do a swap in our bits for i and j (0 and
104). Since all of our bits are in order we simply switch `S[0]` with `S[104]`
which is 0 for 104.

**Second Iteration**
```
i = 1
j = 104

j = (j + key[i % key_length] + S[i]) % 255
j = (104 + key[1 % 5] + 1) % 255
j = (104 + "e" + 1) % 255
j = 206 % 255
j = 206
```

Like before, we can swap `S[1]` with `S[206]` and move to the next round. The
thing to pay attention to here is that we are actively using the previous
state for each swap. The process continues until we have swapped all values of
our bit pool. The result is:

```
[
  124 227 177 171 214 169 219 105 38 82 199 13 175 40 192 119 145 136 114 17 66
  123 56 170 68 165 158 67 146 95 163 76 178 233 193 247 108 118 71 190 57 153
  180 198 236 129 16 39 187 113 245 186 84 22 166 230 131 29 53 120 226 151 88
  217 142 121 185 215 181 157 51 221 222 126 220 81 116 101 42 8 9 250 23 194 112
  252 49 211 137 246 83 94 69 182 183 4 160 87 74 109 11 152 106 132 239 238 111
  34 18 197 164 33 5 26 65 240 46 0 75 79 231 130 62 30 104 36 195 32 167 73 31
  243 191 54 1 248 25 241 91 41 86 122 7 52 93 205 162 223 154 2 140 212 156 6
  128 103 45 148 207 147 15 110 133 55 204 48 43 196 203 139 27 107 210 14 242
  224 125 60 96 19 47 144 202 174 117 253 141 235 208 206 150 188 77 209 200 90
  63 216 234 24 143 173 20 35 232 44 10 3 72 99 78 70 155 201 21 102 58 159 213
  37 218 254 255 161 127 168 179 189 251 80 244 184 237 64 115 89 100 28 59 249
  176 172 135 134 97 12 50 92 229 85 228 61 149 138 225 98
]
```

We now move into the pseudorandom generator portion. Since our plaintext is
14 characters long then we will need a keystream of 14 bits to encrypt with.

**First Iteration**

Our algorithm looks like the following:

```
i = (i + 1) % 255
j = (j + S[i]) % 255

S[i], S[j] = S[j], S[i]

S[(S[i] + S[j]) % 255]
```

Now filling in those solved values we get:

```
i = 1
j = 227

S[1], S[227] = S[227], S[1]

S[(189 + 227) % 255]
S[161]
110
```

Now we see the first bit of the keystream is 110. If we want to get the first
bit of the ciphertext then we convert the first bit of the plaintext to its
ordinal value (65) and XOR it against the first bit of the keystream (110)
giving us 47! This is the first bit of the ciphertext.

We do the same process until our plaintext has been encrypted. Our keysteam
ends up being:

```
[110, 116, 54, 23, 142, 168, 123, 4, 93, 199, 230, 53, 207, 136]
```

Our ciphertext is:

```
[47, 0, 66, 118, 237, 195, 91, 101, 41, 231, 130, 84, 184, 230]
```

Let's test that we're doing this right. With XOR we can encrypt the plaintext
but we can also decrypt it. If we XOR each bit of the keystream against the
corresponding bit of the ciphertext we can see if the decryption works:

```
47  ^ 110 = 65  = A
0   ^ 116 = 116 = t
66  ^ 54  = 116 = t
118 ^ 23  = 97  = a
237 ^ 142 = 99  = c
195 ^ 168 = 107 = k
91  ^ 123 = 32  = <SPACE>
101 ^ 4   = 97  = a
41  ^ 93  = 116 = t
231 ^ 199 = 32  = <SPACE>
130 ^ 230 = 100 = d
84  ^ 53  = 97  = a
184 ^ 207 = 119 = w
230 ^ 136 = 110 = n
```

Looks like it worked. Hopefully that clarifies exactly how it works.

## More Information

While this algorithm has been exploited in a number of ways as not truly secure
(WEP is based on this) it does have some upside. The speed and simplicity of it
makes it very favorable in a number of applications. It can be made efficient
in both hardware and software.

## Cipher Detail

| Details          | Notes            |
| --------- |-------------|
| **Key Sizes** | 40 - 2048 bits |
| **State Sizes** | 2064 bits (1684 Effective) |
| **Rounds** | 1 |
| **Speed** | 7 Cycles per byte on original Pentium |
