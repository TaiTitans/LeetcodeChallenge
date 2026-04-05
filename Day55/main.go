package main

import "sort"

func minimumDifference(nums []int, k int) int {
	length := len(nums)
	if length == 1 {
		return 0
	}

	sort.Ints(nums)
	println("nums: ", nums)
	minDiff := nums[k-1] - nums[0]
	println("minDiff: ", minDiff)
	for i := 1; i < length-k+1; i++ {
		println("nums[i+k-1]: ", nums[i+k-1], "nums[i]: ", nums[i])
		diff := nums[i+k-1] - nums[i]
		println("diff: ", diff)
		if diff < minDiff {
			minDiff = diff
			println("minDiff: ", minDiff)
			println("--------------------------------")
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
