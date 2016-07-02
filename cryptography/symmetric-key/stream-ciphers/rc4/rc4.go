package rc4

// CipherState holds the information about the RC4 Instance
type CipherState struct {
	// S is a permutation of 256 possible bytes
	S [256]int
	// Index pointers for indexing the possible bytes
	i, j uint8
}

// MakeRC4 Returns a new RC4 Instance with the key initiated for excryption
func MakeRC4(key []byte) *CipherState {
	var c CipherState

	l := len(key)

	// First we build the possible bytes with every number in range 0..256
	for i := 0; i < 256; i++ {
		c.S[i] = i
	}

	// This is known as key-scheduling where we take the key and over 256
	// iterations we swap bytes to reach a pseudorandom shuffle of bytes.
	var j uint8 = 0
	for i := 0; i < 256; i++ {
		j = (j + uint8(key[i%l]) + uint8(c.S[i])) % 255
		c.S[i], c.S[j] = c.S[j], c.S[i]
	}
	return &c
}

// Encrypt takes in a sequence of bytes and returns a ciphertext of said bytes
func (c *CipherState) Encrypt(src []byte) []byte {
	result := make([]byte, len(src))
	i, j := c.i, c.j

	// We loop for the length of the source bytes to get xor each byte
	for k, v := range src {
		i = (i + 1) % 255
		j = (j + uint8(c.S[i])) % 255
		c.S[i], c.S[j] = c.S[j], c.S[i]

		result[k] = v ^ uint8(c.S[(c.S[i]+c.S[j])%255])
	}
	c.i, c.j = i, j
	return result
}
