package suffix_array

import "sort"

// Suffix represents a single suffix item which holds a reference to the
// initial string index and the last range of sort indices.
type Suffix struct {
	nr  []int
	idx int
}

// SuffixResult is the group of Suffix objects
type SuffixResult []*Suffix

// ToSuffixArray gets the sorted suffix array result.
func (s SuffixResult) ToSuffixArray() []int {
	var suffixArray []int
	for _, e := range s {
		suffixArray = append(suffixArray, e.idx)
	}
	return suffixArray
}

// Search finds a given pattern within a string using a suffix array. It
// returns two integers, the start index of the suffix array for the match and
// the end index for the suffix array for the match.
func (s SuffixResult) Search(str, pat string) (int, int) {
	suffixArray := s.ToSuffixArray()
	n := len(suffixArray)
	start := 0
	end := n

	for start < end {
		mid := (start + end) / 2
		if pat > str[suffixArray[mid]:n] {
			start = mid + 1
		} else {
			end = mid
		}
	}

	suffixStart := start
	end = n

	for start < end {
		mid := (start + end) / 2
		if pat < str[suffixArray[mid]:n] {
			end = mid
		} else {
			start = mid + 1
		}
	}
	return suffixStart, end
}

// BySuffix is a convenience type to handle sorting by the suffixes.
type BySuffix []*Suffix

// Len satisfies the sort interface.
func (s BySuffix) Len() int {
	return len(s)
}

// Swap satisfies the sort interface.
func (s BySuffix) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Less satisfies the sort interface. In this case, we check to see if the two
// sort indices are equal. If they are then we check which of the second
// indices is greater. If they aren't equal then we find out which of the
// first indices is greater.
func (s BySuffix) Less(i, j int) bool {
	if s[i].nr[0] == s[j].nr[0] {
		return s[i].nr[1] < s[j].nr[1]
	}

	return s[i].nr[0] < s[j].nr[0]
}

// ConstructSuffixArray builds the initial suffix array returning the subsequent suffix array
// along with the sort index matrix which can be used to find the LCP.
func ConstructSuffixArray(s string) (SuffixResult, [][]int) {
	var sortIndex [][]int

	n := len(s)
	suffixes := make(SuffixResult, n)

	// Build the initial sort index row containing integers based on distance from `a`
	sortIndex = append(sortIndex, make([]int, n))
	for i, char := range s {
		sortIndex[0][i] = int(char - 'a')
	}

	// Here we are sorting suffixes in intervals of 2 characters. In this way,
	// we first look at the suffixes as the first 2 characters, first 4, first
	// 8, etc. This allows us to progressively make sorting decisions.
	for step, count := 1, 1; count < n; step, count = step+1, count*2 {
		// Build the array of suffixes where the range is based on the
		// previous batches sort index results. Since we want to store the
		// current sort index and the previous sort index (nr) then we check
		// that we can. If not, we denote it with then -1.
		for i := 0; i < n; i++ {
			nr := make([]int, 2)
			nr[0] = sortIndex[step-1][i]
			if i+count < n {
				nr[1] = sortIndex[step-1][i+count]
			} else {
				nr[1] = -1
			}
			suffixes[i] = &Suffix{nr: nr, idx: i}
		}

		// Sort the suffixes to their order based on the sort indexes collected above.
		sort.Sort(BySuffix(suffixes))

		// Build the next row of sort indexes based on the previous row.
		sortIndex = append(sortIndex, make([]int, n))
		for i := 0; i < n; i++ {
			// This step really checks to see if the current suffix and the
			// one before it are the same substring as one another. If they
			// are, then we want to prior suffix's index to move into the next
			// round.
			if i > 0 && suffixes[i].nr[0] == suffixes[i-1].nr[0] && suffixes[i].nr[1] == suffixes[i-1].nr[1] {
				sortIndex[step][suffixes[i].idx] = sortIndex[step][suffixes[i-1].idx]
			} else {
				sortIndex[step][suffixes[i].idx] = i
			}
		}
	}

	return suffixes, sortIndex
}
