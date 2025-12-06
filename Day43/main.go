package main

func intersect(nums1 []int, nums2 []int) []int {
	counts := make(map[int]int)
	for _, num := range nums1 {
		counts[num]++
	}
	result := []int{}
	for _, num := range nums2 {
		if counts[num] > 0 {
			result = append(result, num)
			counts[num]--
		}
	}
	return result
}

func main() {
	// Example usage:
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2}
	result := intersect(nums1, nums2)
	for _, v := range result {
		print(v, " ")
	}
}
