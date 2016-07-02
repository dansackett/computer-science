package rc4

import "testing"

func TestAlg(t *testing.T) {
	var secret, plaintext, encryptedOutput []byte

	secret = []byte("Secret")
	plaintext = []byte("Attack at dawn")
	encryptedOutput = []byte{135, 128, 194, 160, 112, 180, 75, 155, 220, 52, 116, 4, 157, 192}

	cipher := MakeRC4(secret)

	for i, b := range cipher.Encrypt(plaintext) {
		if b != encryptedOutput[i] {
			t.Errorf("Ciphertext bit at %d (%d) does not match expected bit %d", i, b, encryptedOutput[i])
		}
	}
}
