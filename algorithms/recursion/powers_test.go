package recursion

import "testing"

func TestCanComputeZeroPower(t *testing.T) {
	result := Power(9, 0)

	if result != 1 {
		t.Errorf("9^0 should equal 1, found %d", result)
	}
}

func TestCanComputeOddPower(t *testing.T) {
	result := Power(9, 3)

	if result != 729 {
		t.Errorf("9^3 should equal 729, found %d", result)
	}
}

func TestCanComputeEvenPower(t *testing.T) {
	result := Power(9, 2)

	if result != 81 {
		t.Errorf("9^2 should equal 81, found %d", result)
	}
}

// This test shows that since I'm using integers, we won't get a correct
// answer here since we chop the decimal places.
func TestCanComputeNegativePower(t *testing.T) {
	result := Power(9, -4)

	if result != 0 {
		t.Errorf("9^-4 should equal 0, found %d", result)
	}
}
