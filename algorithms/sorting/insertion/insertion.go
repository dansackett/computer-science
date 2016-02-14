package sort

// InsertionSort takes a slice of integers and sorts it in place.
func InsertionSort(data []int) {
	for i := 1; i < len(data); i++ {
		Insert(data, i-1, data[i])
	}
}

// Insert loops backwards through and array until it finds a value less than
// itself and places itself directly after it.
func Insert(data []int, rightBound, val int) {
	var i int
	for i = rightBound; i >= 0 && data[i] > val; i-- {
		data[i+1] = data[i]
	}
	data[i+1] = val
}
