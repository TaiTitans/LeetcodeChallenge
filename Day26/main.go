package day26

// 3318. Find X-Sum of All K-Long Subarrays I
func xSum(arr []int, k int) int64 {
	n := len(arr)
	if k > n {
		return 0
	}
	var totalSum int64
	var windowSum int64
	for i := 0; i < k; i++ {
		windowSum += int64(arr[i])
	}
	totalSum += windowSum
	for i := k; i < n; i++ {
		windowSum += int64(arr[i]) - int64(arr[i-k])
		totalSum += windowSum
	}
	return totalSum
}
func findXSum(nums []int, k int, x int) []int {
	return []int{}
}
