package main

import (
	"fmt"
	"math"
	"sort"
)

// 3397. Maximum Number of Distinct Elements After Operations

// Solution
//

func maxDistinctElements(nums []int, k int) int {
	sort.Ints(nums)
	nextA := math.MinInt
	distinct := 0
	for _, num := range nums {
		if nextA < num-k {
			nextA = num - k
		}

		if nextA <= num+k {
			distinct++
			nextA++
		}

	}
	return distinct
}

func main() {
	// Example usage
	nums := []int{1, 2, 2, 3, 4}
	k := 2
	result := maxDistinctElements(nums, k)
	fmt.Println(result)
}
