package main

func maxProfit(prices []int) int {

	minPrice := prices[0]
	maxProfit := 0
	for _, v := range prices {
		if v < minPrice {
			minPrice = v
		} else if maxProfit < (v - minPrice) {
			maxProfit = v - minPrice
		}
	}

	return maxProfit
}

//rememeber
