package main

func minMoves(nums []int) int {
	minElement := minInArr(nums)
	result := 0
	for i := 0; i < len(nums); i++ {
		result += nums[i] - minElement
	}
	return result
}

func minInArr(nums []int) int {
	result := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] < result {
			result = nums[i]
		}
	}
	return result
}

func main() {
	nums := []int{1, 2, 3}
	println(minMoves(nums))
}
