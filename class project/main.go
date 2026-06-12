package main

import (	
	"example.com/price-calculator/prices"
)

func main() {
	taxRate := []float64{0.07, 0.08, 0.06, 0.09, 0.05}

	for _, taxRate := range taxRate {
		priceJob := prices.NewTaxIncludedPricesJob(taxRate)
		priceJob.Calculate()
	}
	
}
