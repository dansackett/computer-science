package sort

import "testing"

func TestSelectionSortWithOneValue(t *testing.T) {
	testCase := []int{1}
	target := []int{1}

	SelectionSort(testCase)

	for i, val := range testCase {
		if val != target[i] {
			t.Errorf("Slice values don't match, testCase is %d target is %d", testCase[i], target[i])
		}
	}
}

func TestSelectionSortWithMultipleValues(t *testing.T) {
	testCase := []int{3, 6, 2, 9, 4, 10}
	target := []int{2, 3, 4, 6, 9, 10}

	SelectionSort(testCase)

	for i, val := range testCase {
		if val != target[i] {
			t.Errorf("Slice values don't match, testCase is %d target is %d", testCase[i], target[i])
		}
	}
}
