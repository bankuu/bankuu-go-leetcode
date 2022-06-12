package main

/*

 * 332. Coin Change
 * Submit Count : 12 times
 * Time used : 1.30 Hours

 * Solve by バンクー
 * Code with L#v

 */

import "sort"

func coinChange(coins []int, amount int) int {

	if amount == 0 {
		return 0
	}

	// if amount is odd and all coin is even
	isAmountOdd := amount%2 != 0
	isAllCoinIsEven := true

	for _, coin := range coins {
		if coin%2 != 0 {
			isAllCoinIsEven = false
			break
		}
	}
	if isAmountOdd == true && isAllCoinIsEven == true {
		return -1
	}
	// --

	//sort coin
	sort.Sort(sort.Reverse(sort.IntSlice(coins)))

	// build stack coin
	listSlotChange := map[int]interface{}{}

	// create root coins
	// initial coin
	for _, coin := range coins {
		listSlotChange[coin] = nil
	}

	stackCoin := 1

	for {
		// check stack
		newSlotChange := map[int]interface{}{}
		for slot := range listSlotChange {
			for _, coin := range coins {
				if slot == amount {
					return stackCoin
				} else if slot+coin == amount {
					return stackCoin + 1
				} else {
					if slot+coin <= amount {
						newSlotChange[slot+coin] = nil
					}
				}
			}
		}
		listSlotChange = newSlotChange

		// not found
		if len(newSlotChange) == 0 {
			return -1
		}

		// add stack coin
		stackCoin++

	}
}

func main() {
	result := coinChange([]int{1}, 2)
	print(result)
}
