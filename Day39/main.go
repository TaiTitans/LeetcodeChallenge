package main

// 3512. Minimum Operations to Make Array Sum Divisible by K

func minOperations(nums []int, k int) int {
	totalSum := sumSlice(nums)
	rem := totalSum % k
	if rem == 0 {
		return 0
	}
	return rem
}

func sumSlice(nums []int) int {
	totalSum := 0
	for _, sum := range nums {
		totalSum += sum
	}
	return totalSum
}

func main() {
	// Example usage
	nums := []int{3, 9, 7}
	k := 5
	result := minOperations(nums, k) // Expected output: 4
	println(result)                  // Output the result
}
