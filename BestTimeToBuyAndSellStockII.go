package main

func maxProfit(prices []int) int {

	if prices == nil || len(prices) < 2 {
		return 0
	}

	maxProfit := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			maxProfit += prices[i] - prices[i-1]
		}
	}

	return maxProfit
}