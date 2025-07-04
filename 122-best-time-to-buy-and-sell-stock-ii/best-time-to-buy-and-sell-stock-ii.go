func maxProfit(prices []int) int {
    minPrice := prices[0]; maxProfit := 0; totalProfit := 0;
    for _, v := range prices {
        if minPrice>v {
            minPrice = v
        }

        if v-minPrice > maxProfit {
            totalProfit += v - minPrice
            maxProfit = 0
            minPrice = v
        }
    }
    return totalProfit
}