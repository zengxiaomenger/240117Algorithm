package main

import "fmt"

type record struct {
	price int
	pre   int
	suf   int
}

func findProfit2(prices []int) int {
	ans := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			ans += prices[i] - prices[i-1]
			println(ans)
		}
	}
	return ans
}
func findProfit(prices []int) int {
	Lis_record := []record{}
	for i := 0; i < len(prices); i++ {
		Lis_record = append(Lis_record, record{
			price: prices[i],
			pre:   0,
			suf:   0,
		})
	}
	Lis_record[0].pre = Lis_record[0].price
	for i := 1; i < len(prices); i++ {
		Lis_record[i].pre = min(Lis_record[i].price, Lis_record[i-1].pre)
	}
	Lis_record[len(prices)-1].suf = Lis_record[len(prices)-1].price
	for i := len(prices) - 2; i >= 0; i-- {
		Lis_record[i].suf = max(Lis_record[i].price, Lis_record[i+1].suf)
	}
	maxx := 0
	for i := 0; i < len(prices); i++ {
		if maxx < Lis_record[i].suf-Lis_record[i].pre {
			maxx = Lis_record[i].suf - Lis_record[i].pre
		}
	}
	// fmt.Println(Lis_record)
	return maxx
}
func main() {
	nums := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(findProfit2(nums))
	return
}
