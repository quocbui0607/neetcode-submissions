func maxProfit(prices []int) int {
    maxProfit := 0
    buy := 0
    sell := 1
    for sell < len(prices) {
        if prices[buy] < prices[sell]  {
            maxProfit = max(maxProfit, prices[sell] - prices[buy])
        } else {
			buy = sell
		}
		sell++
    }
    return maxProfit
}
