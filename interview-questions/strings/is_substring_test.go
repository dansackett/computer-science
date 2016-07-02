package interview

import "testing"

func TestIsSubstring(t *testing.T) {
	x := "waterbottle"
	y := "erbottlewat"

	if !IsSubstring(x, y) {
		t.Errorf("'%s' should be a rotation of '%s'", y, x)
	}

	x = "waterbottl"
	y = "erbottlewat"

	if IsSubstring(x, y) {
		t.Errorf("'%s' should NOT be a rotation of '%s'", y, x)
	}
}
