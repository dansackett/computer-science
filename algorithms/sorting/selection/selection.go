package sort

// SelectionSort takes a slice of integers and sorts it in place.
func SelectionSort(data []int) {
	for i, _ := range data {
		min := MinIndex(data, i)
		Swap(data, i, min)
	}
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
