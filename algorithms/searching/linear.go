// ----------------------------------------------------------------------------
// Algorithms - Linear Search
// ----------------------------------------------------------------------------
//
// This is an example of linear search with integers. Instead of complicating
// this with interfaces and reflection, I find it best to use integers to
// illustrate the concept.
//
// With linear search, we essentially go one by one through a list checking if
// the value is equal. If it is, we stop. If not, we keep looking. So for a
// list like the following, if we're looking for 5 then it would take us 5
// guesses to get there.
//
// [1, 2, 3, 4, 5, 6, 7]
//
// ----------------------------------------------------------------------------
// Time Complexity
// ----------------------------------------------------------------------------
//
// This search is defined as running O(n) since it depends on the size of the
// list we're searching through.
//
// ----------------------------------------------------------------------------
package search

func LinearSearch(haystack []int, needle int) int {
	for i, guess := range haystack {
		if guess == needle {
			return i
		}
	}
	return -1
}
