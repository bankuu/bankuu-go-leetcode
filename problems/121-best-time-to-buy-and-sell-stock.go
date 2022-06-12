package main

/*

 * 121. Best Time to Buy and Sell Stock
 * Submit Count : 1 time
 * Time used : 20 m

 * Solve by バンクー
 * Code with L#v

 */

import "fmt"

func maxProfit(prices []int) int {

	maxProfit := 0
	profitNote := 0

	for idx, _ := range prices {
		price := prices[len(prices)-idx-1]
		if profitNote <= price {
			profitNote = price
		} else {
			profit := profitNote - price
			if maxProfit < profit {
				maxProfit = profit
			}
		}
	}
	return maxProfit
}

func main() {
	result := maxProfit([]int{7, 1, 5, 3, 6, 4})
	fmt.Printf("%v", result)
}
