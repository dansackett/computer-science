package sort

import "testing"

func TestQuickSortWithOneValue(t *testing.T) {
	testCase := []int{1}
	target := []int{1}

	QuickSort(testCase)

	for i, val := range testCase {
		if val != target[i] {
			t.Errorf("Slice values don't match, testCase is %d target is %d", testCase[i], target[i])
		}
	}
}

func TestQuickSortWithMultipleValues(t *testing.T) {
	testCase := []int{3, 6, 2, 9, 4, 10}
	target := []int{2, 3, 4, 6, 9, 10}

	QuickSort(testCase)

	for i, val := range testCase {
		if val != target[i] {
			t.Errorf("Slice values don't match, testCase is %d target is %d", testCase[i], target[i])
		}
	}
}
