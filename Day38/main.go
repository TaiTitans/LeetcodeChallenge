package main

func maxProfit(prices []int) int {
	maxProfit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			maxProfit += prices[i] - prices[i-1]
		}
	}
	return maxProfit
}

func main() {
	// Example usage
	prices := []int{7, 1, 5, 3, 6, 4}
	result := maxProfit(prices)
	println(result) // Output: 7
}
