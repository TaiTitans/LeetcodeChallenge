package main

// Solution 1
func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	mappingS := make(map[byte]byte)
	mappingT := make(map[byte]byte)

	for i := 0; i < len(s); i++ {
		c1, c2 := s[i], t[i]
		if val, ok := mappingS[c1]; ok {
			if val != c2 {
				return false
			}
		} else {
			mappingS[c1] = c2
		}
		if val, ok := mappingT[c2]; ok {
			if val != c1 {
				return false
			}
		} else {
			mappingT[c2] = c1
		}
	}
	return true
}

func main() {
	println(isIsomorphic("egg", "add"))     // true
	println(isIsomorphic("foo", "bar"))     // false
	println(isIsomorphic("paper", "title")) // true
	println(isIsomorphic("ab", "aa"))       // false
}
