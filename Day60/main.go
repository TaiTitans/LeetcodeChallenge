package main

func judgeCicle(moves string) bool {
	mapResult := make(map[byte]int)
	// add UDLR
	for _, m := range moves {
		switch m {
		case 'U':
			mapResult[byte('U')]++
		case 'D':
			mapResult[byte('D')]++
		case 'L':
			mapResult[byte('L')]++
		case 'R':
			mapResult[byte('R')]++

		}
	}
	// If U == D and L == R then we are in the same point
	if mapResult[byte('U')] == mapResult[byte('D')] && mapResult[byte('L')] == mapResult[byte('R')] {
		return true
	}
	return false
}

func main() {
	moves := "UD"
	judgeCicle(moves)
	println(moves)
}
