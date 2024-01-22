package main

import (
	"fmt"
)

// За O(N^2)
/* func maxProfit(prices []int) int {
	dp := make([]int, len(prices))

	for i := len(prices) - 1; i > -1; i-- {
		for j := len(prices) - 1; j > i; j-- {
			if prices[j]-prices[i] > dp[j] {
				dp[j] = prices[j] - prices[i]
			}
		}
	}
	return slices.Max(dp)
} */

// За O(N)
func maxProfit(prices []int) int {

	var max, min1, min2, sign int
	if len(prices) == 1 {
		return 0
	}
	if prices[0] > prices[1] {
		max = prices[0]
		min1 = prices[1]
		min2 = prices[1]
		sign = -1
	} else {
		max = prices[1]
		min1 = prices[0]
		min2 = prices[0]
		sign = 1
	}

	for i := 2; i < len(prices); i++ {
		if prices[i] < min2 {
			min2 = prices[i]
			continue
		}
		if prices[i]-min2 > sign*(max-min1) {
			max = prices[i]
			min1 = min2
			sign = 1
		}
		if sign == -1 && prices[i]-min1 > 0 {
			sign = 1
		}
	}
	if sign*(max-min1) > 0 {
		return max - min1
	} else {
		return 0
	}
}

func main() {
	fmt.Printf("%v\n", maxProfit([]int{7, 1, 5, 3, 6, 4}))
}
