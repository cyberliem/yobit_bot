package main

import (
	"log"

	yobit "github.com/go-cryptoexchange/go-yobit"
)

func main() {
	client := yobit.New("key của mi", "secret của mi")
	obs, err := client.GetOrderBook("eth_btc", 10)
	if err != nil {
		log.Fatal("failed to get orderbooks", err)
	}
	log.Printf("orderbooks is %v", obs)
}
