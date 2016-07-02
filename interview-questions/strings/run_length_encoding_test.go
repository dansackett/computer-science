package interview

import (
	"strings"
	"testing"
)

func TestEncode(t *testing.T) {
	s1 := "abc"
	s1ShouldBe := "abc"
	encoded1 := Encode(s1)

	s2 := "abccc"
	s2ShouldBe := "abc3"
	encoded2 := Encode(s2)

	s3 := "aaaabbcccdeee"
	s3ShouldBe := "a4b2c3de3"
	encoded3 := Encode(s3)

	if s1 != encoded1 {
		t.Errorf("Length of '%s' should be the same as the encoded string", s1)
	}

	if s1ShouldBe != encoded1 {
		t.Errorf("string '%s' is not the same as '%s'", s1, s1ShouldBe)
	}

	if len(s2) == len(encoded2) {
		t.Errorf("Length of encoded string '%s' should not be the same as initial string", s2)
	}

	if s2ShouldBe != encoded2 {
		t.Errorf("string '%s' is not the same as '%s'", s2, s2ShouldBe)
	}

	if len(s3) == len(encoded3) {
		t.Errorf("Length of encoded string '%s' should not be the same as initial string", s3)
	}

	if s3ShouldBe != encoded3 {
		t.Errorf("string '%s' is not the same as '%s'", s3, s3ShouldBe)
	}
}

func TestDecode(t *testing.T) {
	s1 := "abc"
	encoded1 := Encode(s1)
	decoded1 := Decode(encoded1)

	s2 := "abccc"
	encoded2 := Encode(s2)
	decoded2 := Decode(encoded2)

	s3 := "aaaabbcccdeee"
	encoded3 := Encode(s3)
	decoded3 := Decode(encoded3)

	if strings.Compare(s1, decoded1) != 0 {
		t.Errorf("String '%s' should be the same as the decoded string but is '%s'", s1, decoded1)
	}

	if strings.Compare(s2, decoded2) != 0 {
		t.Errorf("String '%s' should be the same as the decoded string but is '%s'", s2, decoded2)
	}

	if strings.Compare(s3, decoded3) != 0 {
		t.Errorf("String '%s' should be the same as the decoded string but is '%s'", s3, decoded3)
	}
}
