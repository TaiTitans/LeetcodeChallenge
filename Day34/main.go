package main

import "sort"

// Set solution
func a(nums []int, original int) int {
	currentValue := original
	numSet := make(map[int]struct{})
	for _, num := range nums {
		numSet[num] = struct{}{}
	}
	for {
		if _, exists := numSet[currentValue]; exists {
			currentValue *= 2
		} else {
			break
		}
	}
	return currentValue
}

// Sorting solution
func b(nums []int, original int) int {
	currentValue := original
	// Sort nums
	sort.Ints(nums)
	numsLen := len(nums)
	for i := 0; i < numsLen; i++ {
		if nums[i] == currentValue {
			currentValue *= 2
		} else if nums[i] > currentValue {
			break
		}
	}
	return currentValue
}

func main() {
	nums := []int{8, 19, 4, 2, 15, 3}
	original := 2
	result := a(nums, original)
	println(result) // Expected output: 16
	result2 := b(nums, original)
	println(result2) // Expected output: 16
}
