package main

import "sort"

func minPairSum(nums []int) int {
	sort.Ints(nums)
	length := len(nums)
	maxPairSum := 0
	for i := 0;i < length;i++ {
		currentSum := nums[i] + nums[length - 1 - i]
		if currentSum > maxPairSum {
			maxPairSum = currentSum
		}	
	}
	return maxPairSum
}

func main() {
	nums := []int{3,5,4,2,4,6} // sort: [2,3,4,4,5,6] -> (2+6), (3+5), (4+4) -> 8, 8, 8
	result := minPairSum(nums)
	println(result) // Expected output: 8
}
