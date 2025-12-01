package main

func firstUniqChar(s string) int {
	charCount := make(map[rune]int)
	for _, char := range s {
		charCount[char]++
	}
	for i, char := range s {
		if charCount[char] == 1 {
			return i
		}
	}
	return -1

}

func main() {
	s := "leetcode"
	result := firstUniqChar(s)
	println(result) // Output the result is 0
}
