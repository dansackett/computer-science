// ----------------------------------------------------------------------------
// Algorithms - Merge Sort
// ----------------------------------------------------------------------------
//
// Merge sort is a recursive sorting algorithm which utilizes the divide and
// conquer concept. It works by taking a set of data, splitting it into two
// halves, and merging those two halves in sorted order. For example:
//
// If we have a set: [3, 1, 6, 0, 5]
//
// We split that into two halves [3, 1, 6] and [0, 5]. We then split those
// halves into halves again: [3, 1] and [6] and [0] and [5]. We split the last
// half into two and then we begin the merge as we bubble back up:
//
// On one side we get back [1, 3] and [6]
// On the other side we get back [0, 5]
//
// Into the next step we get [1, 3, 6] and [0, 5].
//
// In the final merge at the top, we get [0, 1, 3, 5, 6] giving us the sorted
// set.
//
// The merge technique used here checks if both the left and right sets have
// values in it. If they do, it finds which element at the 0 index of each set
// is smaller. It adds that value to the resulting set and runs the check
// again. Since it is common then to have one set be finished while the other
// is not, we do a final step of appending the left or right values onto the
// end of our resulting set to get the full set back sorted.
//
// ----------------------------------------------------------------------------
// Time Complexity
// ----------------------------------------------------------------------------
//
// Worst case: Θ(n lg n)
// Best case: Θ(n lg n)
// Average case: Θ(n lg n)
//
// ----------------------------------------------------------------------------
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
