package main

import "fmt"

func main() {
	 prices  := []float64{10.5, 20.0, 15.75, 30.0, 25.5}
	 taxRate := []float64{0.07, 0.08, 0.06, 0.09, 0.05}

	 result := make(map[float64] []float64)

	 for _, taxRate := range taxRate {
		taxIncludedPrice := make([]float64, len(prices))
		for priceIndex, price := range prices {
			taxIncludedPrice[priceIndex] = price * (1 + taxRate)
		}
		result[taxRate] = taxIncludedPrice
	 }

	 fmt.Println(result)
}