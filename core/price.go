package core

import "fmt"

var (
	InitCore             = initCore
	SetCurrentPairPrice  = setCurrentPairPrice
	GetCurrenctPairPrice = getCurrenctPairPrice
)

var prices map[string]int64

func initCore() {
	prices = make(map[string]int64)
}
func setCurrentPairPrice(pair string, price int64) {
	fmt.Println("set price to ", price)
	prices[pair] = price
}

func getCurrenctPairPrice(pair string) int64 {
	price := prices[pair]
	return price
}
