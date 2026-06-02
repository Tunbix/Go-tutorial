package main

import "fmt"

func main()  {
	var productNames [4]string = [4]string{"Book", "Pen", "Pencil", "Eraser"}
	prices := [4]float64{1.99, 2.99, 3.99, 10.99}
	fmt.Println(prices)
	fmt.Println(productNames)

	fmt.Println(prices[2])

	featuredPrices := prices[1:]
	highlightedPrices := prices[:1]
	fmt.Println(featuredPrices)
	fmt.Println(highlightedPrices)
	fmt.Println(prices)
	fmt.Println(len(featuredPrices), cap(featuredPrices))

	highlightedPrices = highlightedPrices[:3]
	fmt.Println(highlightedPrices)
	fmt.Println(len(featuredPrices), cap(featuredPrices))

}