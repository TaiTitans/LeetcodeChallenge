package main

// 3354. Make Array Elements Equal to Zero
func countValidSelections(nums []int) int {
	s := 0
	for _, num := range nums {
		s += num
	}
	ans, l := 0, 0
	for _, x := range nums {
		if x == 0 {
			if 2*l == s {
				ans += 2
			} else if abs(2*l-s) == 1 {
				ans += 1
			}
		}
		l += x
	}
	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
func main() {
	// Example usage
	nums := []int{1, 0, 2, 0, 3}
	result := countValidSelections(nums)
	println(result) // Output: 2
}
