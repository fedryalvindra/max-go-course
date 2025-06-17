package main

import (
	"fmt"
)

type Product struct {
	title string
	id    string
	price float64
}

func main() {
	prices := []float64{10.99, 8.99}
	fmt.Println(prices[0:1])
	prices[1] = 9.99
	
	prices = append(prices, 5.99)
	prices = prices[1:]
	fmt.Println(prices)
}

// func main() {
// 	var productNames [4]string = [4]string{"A book"}

// 	productNames[2] = "A Carpet"

// 	prices := [4]float64{
// 		0.1,
// 		0.2,
// 		0.3,
// 		0.4,
// 	}

// 	fmt.Println(prices)
// 	fmt.Println(productNames)

// 	fmt.Println(prices[2])

// 	// 1 included 3 excluded
// 	featuredPrices := prices[1:]
// 	featuredPrices[0] = 0.199
// 	highlightedPrices := featuredPrices[:1]
// 	fmt.Println(highlightedPrices)
// 	fmt.Println(prices)
// 	fmt.Println(len(highlightedPrices), cap(highlightedPrices))

// 	highlightedPrices = highlightedPrices[:3]
// 	fmt.Println(highlightedPrices)
// 	fmt.Println(len(highlightedPrices), cap(highlightedPrices))
// }
