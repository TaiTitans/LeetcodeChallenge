package main

import "strings"

func countPalindromicSubsequence(s string) int {
	result := 0
	// Loop characters from 'a' to 'z'
	for ch := 'a'; ch <= 'z'; ch++ {
		first := strings.Index(s, string(ch))
		last := strings.LastIndex(s, string(ch))

		println("Character:", string(ch), "First:", first, "Last:", last)

		// Find the first and last occurrence of the character
		if first != -1 && last != -1 && last-first > 1 {
			// Set to track unique middle characters
			seen := make(map[byte]bool)
			for i := first + 1; i < last; i++ {
				seen[s[i]] = true
			}
			result += len(seen)
			println("Result: ", result)
		}
	}
	return result
}

func main() {
	// Example usage
	s := "bbcbaba"
	println(countPalindromicSubsequence(s)) // Output: 3
}
