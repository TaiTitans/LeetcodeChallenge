package main

import "sort"

type Item struct {
	Index int
	Val  int
}
func a(arr []int) []int {
	items := make([]Item, len(arr))
	for i, v := range arr {
		items[i] = Item{Index: i, Val: v}
	}

	sort.Slice(items, func(x, y int) bool {
		return items[x].Val < items[y].Val
	})
	result := make([]int, len(arr))
	rank := 0
	prevVal := 0
	for i, item := range items {
		if i == 0 || item.Val != prevVal {
			rank++
			prevVal = item.Val
		}
		result[item.Index] = rank
	}
	return result
}