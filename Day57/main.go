package main

import "sort"

func topKFrequent(nums []int, k int) []int {
	mapTemp := make(map[int]int)
	for _, num := range nums {
		mapTemp[num] = mapTemp[num] + 1
	}
	sortedMap := make([]int, 0)
	for key := range mapTemp {
		sortedMap = append(sortedMap, key)
	}
	sort.Ints(sortedMap)

	return sortedMap[:k]
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	result := topKFrequent(nums, k)
	println(result)
}
