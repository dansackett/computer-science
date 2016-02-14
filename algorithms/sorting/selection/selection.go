package sort

// SelectionSort takes a slice of integers and sorts it in place.
func SelectionSort(data []int) {
	for i, _ := range data {
		min := MinIndex(data, i)
		Swap(data, i, min)
	}
}

// Swap takes two indices and switches the values between them.
func Swap(data []int, i, j int) {
	data[i], data[j] = data[j], data[i]
}

// MinIndex finds the index of the next smallest number in a slice.
func MinIndex(data []int, initial int) int {
	minValue := data[initial]
	minIndex := initial

	for i := initial; i < len(data); i++ {
		if data[i] < minValue {
			minIndex = i
			minValue = data[i]
		}
	}

	return minIndex
}
