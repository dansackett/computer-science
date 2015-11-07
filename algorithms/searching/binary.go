// ----------------------------------------------------------------------------
// Algorithms - Binary Search
// ----------------------------------------------------------------------------
//
// This is an example of binary search with integers. Instead of complicating
// this with interfaces and reflection, I find it best to use integers to
// illustrate the concept.
//
// With Binary Search, the goal is to constantly work on smaller data sets in
// order to limit the number of guesses one must make. Based on the size of
// the array (slice) we're searching, we can quickly find an answer by
// guessing the number in the middle of our maximum and minimum numbers. So
// for instance, if we have:
//
// [1, 2, 3, 4, 5, 6, 7]
//
// And the number we're looking for is 5 then in a linear search it would
// take us 5 guesses to get there. With binary search, it would only take us 3
// guesses:
//
// First Guess: Index at 3 --- Since our number is higher than 4, we discard
// the first half and our new set becomes [5, 6, 7].
//
// Second Guess: Index at 5 --- Since our number is less than 6, we discard
// the upper half of the set leaving us with only one more number in the set
// [5].
//
// Third Guess: Index at 4 --- This matches our number and we are done.
//
// ----------------------------------------------------------------------------
// Time Complexity
// ----------------------------------------------------------------------------
//
// This search is defined as running O(lg n) as it exponentially gets better.
//
// ----------------------------------------------------------------------------
package search

func BinarySearch(haystack []int, needle int) int {
	min := 0
	max := len(haystack) - 1

	for max >= min {
		guess := (max + min) / 2
		if haystack[guess] == needle {
			return guess
		} else if haystack[guess] < needle {
			min = guess + 1
		} else {
			max = guess - 1
		}
	}
	return -1
}
