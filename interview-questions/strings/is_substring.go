package interview

import "strings"

func IsSubstring(s, t string) bool {
	l := len(s)
	if l == len(t) && l > 0 {
		s1s1 := s + s
		return strings.Contains(s1s1, t)
	}
	return false
}
