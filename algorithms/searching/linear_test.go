package search

import "testing"

func TestLinearCanFindInteger(t *testing.T) {
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53,
		59, 61, 67, 71, 73, 79, 83, 89, 97}

	result1 := LinearSearch(primes, 73)
	result2 := LinearSearch(primes, 11)
	result3 := LinearSearch(primes, 65)

	if result1 != 20 {
		t.Errorf("73 is at index 20 not %d", result1)
	}

	if result2 != 4 {
		t.Errorf("11 is at index 4 not %d", result2)
	}

	if result3 != -1 {
		t.Errorf("65 IS NOT in this array, found it at %d", result3)
	}
}
