package rc4

import (
	"fmt"
	"testing"
)

func TestAlg(t *testing.T) {
	cipher := MakeRC4([]byte("pwd12"))
	result := cipher.Encrypt([]byte("Math 310 Proves!"))
	fmt.Println(result)
	for _, b := range result {
		fmt.Printf("%X", b)
	}
	fmt.Println()
}
