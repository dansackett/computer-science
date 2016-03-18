package sort

// InsertionSort is the internal swapping algorithm.
func InsertionSort(data []int, interval int) {
	// Outer is the right bound
	for outer := interval; outer < len(data); outer++ {
		val := data[outer]
		inner := outer

		// Inner is the left bound
		for inner > interval-1 && data[inner-interval] >= val {
			Swap(data, inner, inner-1)
			inner -= interval
		}

		data[inner] = val
	}
}

// ShellSort takes a slice of integers and sorts it in place.
func ShellSort(data []int) {
	interval := (len(data) / 2) + 1

	for interval > 0 {
		InsertionSort(data, interval)
		interval = (interval - 1) / 3
	}
}

// Swap takes two indices and switches the values between them.
func Swap(data []int, i, j int) {
	data[i], data[j] = data[j], data[i]
}
