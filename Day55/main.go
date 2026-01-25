package main

import "sort"

func minimumDifference(nums []int, k int) int {
	length := len(nums)
	if length == 1 {
		return 0
	}

	sort.Ints(nums)
	minDiff := nums[k-1] - nums[0]
	for i := 1; i < length-k+1; i++ {
		diff := nums[i+k-1] - nums[i]
		if diff < minDiff {
			minDiff = diff
		}
	}
	return minDiff

}

func main() {
	nums := []int{9, 4, 1, 7}
	k := 2
	result := minimumDifference(nums, k)
	println(result) // Expected output: 2
}
