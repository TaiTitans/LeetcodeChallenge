package main

import "sort"

func minMoves2(nums []int) int {
	sort.Ints(nums)
	midElement := nums[len(nums)/2]
	result := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > midElement {
			result += nums[i] - midElement
		} else {
			result += midElement - nums[i]
		}
	}
	return result
}
