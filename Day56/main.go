package main
import "sort"

func minimumCost(nums []int) int {
    if len(nums) == 3 {
		return nums[0] + nums[1] + nums[2]
	}
	totalCost := 0
	// Pick 3 elemnt from array [first element, (pick 2 elements with smallest value)] -> result += first element + smallest + second smallest (not sort)
	firstElement := nums[0]
	// Find two smallest elements in the array
	// Temp array append all elements except firstElement
	temp := make([]int, 0, len(nums)-1)	
	for i := 1; i < len(nums); i++ {
		temp = append(temp, nums[i])
	}
	sort.Ints(temp)
	smallest := temp[0]
	secondSmallest := temp[1]
	totalCost += firstElement + smallest + secondSmallest
	return totalCost
}

func main() {

	nums := []int{10,3,1,1}
	result := minimumCost(nums)
	println(result) // Expected output: 12
}