package suffix_array

import "testing"

func TestSuffixArrayConstruction(t *testing.T) {
	str := "banana"
	sa, _ := ConstructSuffixArray(str)

	if len(sa) != 6 {
		t.Errorf("Suffix array was not contructed properly")
	}

	res := []int{5, 3, 1, 0, 4, 2}
	for i, a := range sa.ToSuffixArray() {
		if a != res[i] {
			t.Errorf("Suffix array was not sorted properly", a, "is out of place")
		}
	}
}

func TestSuffixArraySearch(t *testing.T) {
	str := "banana"
	sa, _ := ConstructSuffixArray(str)
	s, _ := sa.Search(str, "na")

	if s != 4 {
		t.Errorf("Suffix Array search did not find `na` at index 4")
	}
}
