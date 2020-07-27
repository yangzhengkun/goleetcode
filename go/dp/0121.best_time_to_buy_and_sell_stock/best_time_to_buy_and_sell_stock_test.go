package _121_best_time_to_buy_and_sell_stock

import (
	"fmt"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	var prices = []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))

	prices = []int{7, 6, 4, 3, 1}
	fmt.Println(maxProfit(prices))
}
