package sort

// QuickSort is the entry to the sorter taking in just the data to be sorted
// determining the high and low values for us.
func QuickSort(data []int) {
	quickSort(data, 0, len(data)-1)
}

// quickSort is the recursive function which breaks the data into smaller
// chunks based on a calculated pivot value.
func quickSort(data []int, low, high int) {
	if low < high {
		pivot := Partition(data, low, high)
		quickSort(data, low, pivot-1)
		quickSort(data, pivot+1, high)
	}
}

// Swap takes two indices and switches the values between them.
func Swap(data []int, i, j int) {
	data[i], data[j] = data[j], data[i]
}

// Partition swaps values in the data while holding a reference to the pivot
// value to return for the recursive search to decrease.
func Partition(data []int, low, high int) int {
	pivot := data[high]
	pivotIndex := low
	for j := low; j <= high-1; j++ {
		if data[j] <= pivot {
			Swap(data, pivotIndex, j)
			pivotIndex++
		}
	}
	Swap(data, pivotIndex, high)
	return pivotIndex
}
