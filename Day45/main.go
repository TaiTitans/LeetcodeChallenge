package main

func intSqrt(n int) int {
	if n == 0 {
		return 0
	}
	x := n
	y := (x + 1) / 2
	for y < x {
		x = y
		y = (x + n/x) / 2
	}
	return x
}

func countTriples(n int) int {
	count := 0
	for a := 1; a <= n; a++ {
		for b := 1; b <= n; b++ {
			cSquared := a*a + b*b
			c := intSqrt(cSquared)
			if c*c == cSquared && c <= n {
				count++
			}
		}
	}
	return count
}

func main() {
	// Example usage
	result := countTriples(5)
	println(result) // Output: 2
}
