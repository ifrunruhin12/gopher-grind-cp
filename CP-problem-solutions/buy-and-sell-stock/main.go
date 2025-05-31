//https://leetcode.com/problems/best-time-to-buy-and-sell-stock/description/

//Dynamic Programming

package main

import "fmt"

func maxTwo(a, b int) int {
    if a > b {
        return a
    }else {
        return b
    }
}

func maxProfit(prices []int) int {
    currentValue := prices[0]
    currentProfit := 0
	for _, val := range prices{
        if currentValue > val {
            currentValue = val
        }
        currentProfit = maxTwo(currentProfit, val - currentValue)
    }
    return currentProfit
}


func main() {
	s := []int{7,1,5,3,6,4}
	res := maxProfit(s)
	fmt.Println(res)
}
