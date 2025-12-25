package main

func minimumBoxes(a []int, capacity []int) int {
	sumA := 0
	for _, v := range a {
		sumA += v
	}

	count := 0

	// sort capacity in descending order
	for i := 0; i < len(capacity)-1; i++ {
		for j := 0; j < len(capacity)-i-1; j++ {
			if capacity[j] < capacity[j+1] {
				capacity[j], capacity[j+1] = capacity[j+1], capacity[j]
			}
		}
	}

	for _, cap := range capacity {
		sumA -= cap
		count++
		if sumA <= 0 {
			return count
		}
	}
	return count

}

func main() {
	// Example usage
	a := []int{1, 3, 2}
	capacity := []int{4, 3, 1, 5, 2}
	result := minimumBoxes(a, capacity)
	println(result) // Output: 2
}
