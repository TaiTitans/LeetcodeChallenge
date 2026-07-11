package main

func longestContinuousSubstring(s string) int {
	resultCount := 1

	for i := 0; i < len(s); i++ {
		count := 1

		for j := i + 1; j < len(s); j++ {
			if s[j]-s[j-1] == 1 {
				count++
			} else {
				break
			}

			if count > resultCount {
				resultCount = count
			}
		}
	}
	return resultCount
    
}