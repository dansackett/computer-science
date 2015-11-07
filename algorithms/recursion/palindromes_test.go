package recursion

import "testing"

func TestSingleLetterIsPalindrome(t *testing.T) {
	result := IsPalindrome("a")

	if !result {
		t.Errorf("`a` should be a palindrome but our function says it is not.")
	}
}

func TestWordsArePalindromes(t *testing.T) {
	result1 := IsPalindrome("brick")
	result2 := IsPalindrome("racecar")

	if result1 {
		t.Errorf("`brick` should not be a palindrome but our function says it is.")
	}

	if !result2 {
		t.Errorf("`racecar` should be a palindrome but our function says it is not.")
	}
}
