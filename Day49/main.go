package main

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	i, j := 0, 0
	count := 0

	for i < len(g) && j < len(s) {
		if s[j] >= g[i] {
			count++
			i++
		}
		j++
	}
	return count
}

func main() {
	// Example usage
	c := []int{1, 2, 3}
	ck := []int{3}
	result := findContentChildren(c, ck)
	println(result) // Output: 1
}
