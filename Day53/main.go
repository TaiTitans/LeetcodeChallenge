package main

func sumFourDivisors(nums []int) int {
	totalSum := 0
	for _, num := range nums {
		divisors := []int{}
		for i := 1; i*i <= num; i++ {
			if num%i == 0 {
				divisors = append(divisors, i)
				if i != num/i {
					divisors = append(divisors, num/i)
				}
			}
			if len(divisors) > 4 {
				break
			}
		}
		if len(divisors) == 4 {
			for _, d := range divisors {
				totalSum += d
			}
		}
	}
	return totalSum
}

func main() {
	// Input: nums = [21,4,7]
	// Output: 32
	print(sumFourDivisors([]int{21, 4, 7}))
}
