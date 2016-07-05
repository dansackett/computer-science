package kmp

func InitializeFailureTable(needle string, table *[]int) {
	var curPos, nextCandidate int

	// Dereference table pointer for convenience
	T := (*table)

	curPos, nextCandidate = 2, 0
	T[0], T[1] = -1, 0

	for curPos < len(needle) {
		if needle[curPos-1] == needle[nextCandidate] {
			T[curPos] = nextCandidate + 1
			nextCandidate++
			curPos++
		} else if nextCandidate > 0 {
			nextCandidate = T[nextCandidate]
		} else {
			T[curPos] = 0
			curPos++
		}
	}
}

func KMP(haystack, needle string) int {
	var T []int
	var curMatchIdx, curPatternIdx int

	// Initialize FailureTable
	T = make([]int, len(needle))
	InitializeFailureTable(needle, &T)

	// curMatchIdx is current match start in the haystack
	curMatchIdx = 0

	// curPatternIdx is the current index of the pattern we're looking for
	curPatternIdx = 0

	for curMatchIdx+curPatternIdx < len(haystack) {
		if needle[curPatternIdx] == haystack[curMatchIdx+curPatternIdx] {
			if curPatternIdx == len(needle)-1 {
				return curMatchIdx
			}
			curPatternIdx++
		} else {
			if T[curPatternIdx] > -1 {
				curMatchIdx = curMatchIdx + curPatternIdx - T[curPatternIdx]
				curPatternIdx = T[curPatternIdx]
			} else {
				curPatternIdx = 0
				curMatchIdx++
			}
		}
	}
	return -1
}
