package sort

// Swap takes two indices and switches the values between them.
func Swap(data []int, i, j int) {
	data[i], data[j] = data[j], data[i]
}
