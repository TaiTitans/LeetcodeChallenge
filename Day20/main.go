package main

func lengthOfLastWord(s string) int {
	length := 0
	// Start from the end and skip trailing spaces
	i := len(s) - 1
	for i >= 0 && s[i] == ' ' {
		i--
	}
	// Count characters until a space or start of string
	for i >= 0 && s[i] != ' ' {
		length++
		i--
	}
	return length
}
