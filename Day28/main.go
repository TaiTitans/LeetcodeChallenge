package day28

func repeatedSubstringPattern(s string) bool {
	if len(s) == 1 {
		return false
	}

	index := (len(s) + 2 - 1) / 2
	for index > 0 {
		substr := s[0:index]
		// Check value substr == [index:end]
		match := true
		for i := index; i < len(s); i += index {
			end := i + index
			if end > len(s) {
				end = len(s)
			}
			if s[i:end] != substr {
				match = false
				break
			}
		}
		if match {
			return true
		}
		index--
	}

	return false
}
