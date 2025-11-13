package main

import "strings"

func reverseWords(s string) string {
	// Convert the string to a slice
	words := []rune(s)
	n := len(words)
	reverse := func(left, right int) {
		for left < right {
			words[left], words[right] = words[right], words[left]
			left++
			right--
		}
	}
	reverse(0, n-1)

	start := 0
	for i := 0; i <= n; i++ {
		if i == n || words[i] == ' ' {
			reverse(start, i-1)
			start = i + 1
		}
	}
	return string(words)
}

func reverseWords2(s string) string {
	// Convert the string to a slice a word is element of slice
	slices := strings.Split(s, " ")
	for _, word := range slices {
		// Reverse each word
		runes := []rune(word)
		left, right := 0, len(runes)-1
		for left < right {
			runes[left], runes[right] = runes[right], runes[left]
			left++
			right--
		}
		// Update the slice with the reversed word
		word = string(runes)
	}
	// Convert the slice back to a string
	return strings.Join(slices, " ")
}
