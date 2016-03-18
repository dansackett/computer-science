package sort

// BubbleSort takes a slice of integers and sorts it in place.
func BubbleSort(data []int) {
	n := len(data)
	swapped := true
	for swapped {
		swapped = false
		for i := 1; i <= n-1; i++ {
			if data[i-1] > data[i] {
				Swap(data, i-1, i)
				swapped = true
			}
		}
	}
}

// Swap takes two indices and switches the values between them.
func Swap(data []int, i, j int) {
	data[i], data[j] = data[j], data[i]
}
