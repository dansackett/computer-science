package sort

// MergeSort takes a slice of integers and sorts it
func MergeSort(data []int) []int {
	if len(data) <= 1 {
		return data
	}

	left := MergeSort(data[:len(data)/2])
	right := MergeSort(data[len(data)/2:])

	return Merge(left, right)
}

// Merge takes the left and right side merging the two into a sorted slice
func Merge(left, right []int) []int {
	var result []int

	// While both the left and right sides have values, sort accordingly
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}

	// Add any extra values from the left and right sides
	result = append(result, left...)
	result = append(result, right...)
	return result
}
