package main

func minArray(numbers []int) int {
	min := numbers[0]
	for _, n := range numbers {
		if n < min {
			min = n
		}
	}
	return min
}

func product(nums []int) int {
	newArr := make([]int, len(nums))
	// convert element to string in nums
	numsStr := make([]string, len(nums))

	for i, n := range nums {
		numsStr[i] = strconv.Itoa(n)
	}

	// sum in element in newArr
	for i, n := range numsStr {
		sum := 0
		for _, c := range n {
			sum += int(c - '0')
		}
		newArr[i] = sum
	}

	result := minArray(newArr)
	return result

}
