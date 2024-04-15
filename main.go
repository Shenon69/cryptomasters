package main

import (
	"fmt"

	"github.com/Shenon69/cryptomasters/api"
)

func main() {
	rate, err := api.GetRate("BTC")

	if err == nil {
		fmt.Printf("The rate for %v is %.2f", rate.Currency, rate.Price)
	}
}