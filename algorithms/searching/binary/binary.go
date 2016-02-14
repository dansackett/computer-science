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
