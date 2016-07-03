package singleton

import "testing"

func TestSingleton(t *testing.T) {
	obj := GetInstance()
	obj2 := GetInstance()

	if obj != obj2 {
		t.Errorf("The singleton instances should be the same object")
	}
}
