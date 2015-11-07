// ----------------------------------------------------------------------------
// Algorithms - Recursive Palindrome
// ----------------------------------------------------------------------------
//
// Palindromes are words that are spelled the same way both forwards and
// backwards. We can use recursion to determine this by checking if the first
// and last letters are equal, stripping the first and last letters, and
// checking again until we are sure that the word is in fact a palindrome.
//
// ----------------------------------------------------------------------------
package recursion

// FirstChar finds the first character of a string
func FirstChar(str string) string {
	return string(str[0])
}

// LastChar finds the last character of a string
func LastChar(str string) string {
	return string(str[len(str)-1])
}

// MiddleChars finds the characters between the first and last characters
func MiddleChars(str string) string {
	return string(str[1 : len(str)-1])
}

// IsPalindrome returns whether a string is a palindrome through recursion
func IsPalindrome(str string) bool {
	if len(str) <= 1 {
		return true
	}

	if FirstChar(str) != LastChar(str) {
		return false
	}

	return IsPalindrome(MiddleChars(str))
}
