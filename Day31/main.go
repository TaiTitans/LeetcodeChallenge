package main

import "fmt"

func subsets(nums []int) [][]int {
	// Initialize the result with the empty subset
	result := [][]int{{}}
	// Iterate through each number in the input slice
	for _, num := range nums {
		// For each existing subset, create a new subset that includes the current number
		n := len(result)
		for i := 0; i < n; i++ {
			// Create a new subset by appending the current number to the existing subset
			newSubset := append([]int{}, result[i]...)
			// Append the current number to the new subset
			newSubset = append(newSubset, num)
			// Add the new subset to the result
			result = append(result, newSubset)
		}
	}
	return result
}

func main() {
	// Example usage
	nums := []int{1, 2, 3}
	subsetsResult := subsets(nums)
	_ = subsetsResult
	fmt.Println(subsetsResult)
}
