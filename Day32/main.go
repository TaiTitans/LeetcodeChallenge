package main

func kLengthApart(nums []int, k int) bool {
	previousOneIndex := -(k + 1) // Initialize to a value that ensures the first '1' is valid
	for i, num := range nums {
		if num == 1 {
			if i-previousOneIndex-1 < k {
				return false
			}
			previousOneIndex = i
		}
	}
	return true
}

// EXAMPLE USAGE
func main() {
	nums := []int{1, 0, 0, 1, 0, 1}
	k := 2
	result := kLengthApart(nums, k)
	_ = result
	println(result)
}
