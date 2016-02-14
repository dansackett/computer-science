package search

func LinearSearch(haystack []int, needle int) int {
	for i, guess := range haystack {
		if guess == needle {
			return i
		}
	}
	return -1
}
