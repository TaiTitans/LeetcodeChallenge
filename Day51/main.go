package main

func plusOne(digits []int) []int {
	resultArray := []int{}

	// convert f to string
	str := ""
	for _, v := range digits {
		str += string(rune(v + '0'))
	}
	// str + 1
	num := 0
	for i := len(str) - 1; i >= 0; i-- {
		num = int(str[i]-'0') + 1
		if num == 10 {
			str = str[:i] + "0" + str[i+1:]
		} else {
			str = str[:i] + string(rune(num+'0')) + str[i+1:]
			break
		}
	}
	// if first char is 0, add 1 at the beginning
	if str[0] == '0' {
		str = "1" + str
	}
	// convert str to array
	for i := 0; i < len(str); i++ {
		resultArray = append(resultArray, int(str[i]-'0'))
	}
	return resultArray
}

func main() {
	// Example usage
	nums := []int{4, 3, 2, 1}
	result := plusOne(nums)
	for _, v := range result {
		print(v)
	}
}
