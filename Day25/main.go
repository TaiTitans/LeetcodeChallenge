package main

// 3370. Smallest Number With All Set Bits
// Hint: Find the strictly greater power of 2, and subtract 1 from it.
func smallestNumber(n int) int {
	p := 1
	for p <= n {
		p <<= 1
	}
	return p - 1
}

func isSetBit(x int) bool {
	return x&(x-1) == 0
}
func main() {
	// Example usage
	n := 5
	result := smallestNumber(n)
	println(result) // Output: 7
}
