// ----------------------------------------------------------------------------
// Algorithms - Recursive Powers
// ----------------------------------------------------------------------------
//
// When computing a number to a specific power, you can use recursion to make
// life easier. We can do this following these rules:
//
// - Anything to the power of 0 == 1
// - Any negative power == 1 / x^-n
// - Any positive even number == x^n/2 * x^n/2
// - Any positive odd number == x^n-1 * x
//
// ----------------------------------------------------------------------------
package recursion

func Power(x, n int) int {
	if n == 0 {
		return 1
	}

	if n < 0 {
		return int(1 / Power(x, -n))
	}

	if n%2 != 0 {
		return Power(x, n-1) * x
	}

	if n%2 == 0 {
		result := Power(x, n/2)
		return result * result
	}

	return 0
}
