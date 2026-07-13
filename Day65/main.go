package main

func sequentialDigits(low int, high int) []int {
	var result []int
	for i := 1; i <= 9; i++ {
		num := i
		for j := i + 1; j <= 9; j++ {
			num = num*10 + j
			if num >= low && num <= high {
				result = append(result, num)
			}
			if num > high {
				break
			}
		}
	}
	sort.Ints(result)
	return result
}