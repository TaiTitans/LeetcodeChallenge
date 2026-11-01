package main

func maximumProduct(nums []int) int {
	result := 0
	if len(nums) == 3 {
		return nums[0] + nums[1] + nums[2]
	}

	sort.Ints(nums)

	start := nums[len(nums)-3]

	for i := len(nums) - 3; i < len(nums)-1; i++ {
		temp := start * nums[i+1]
		start = temp
		result += temp
	}

	return result

}
