package main

import (
	"fmt"

	"example.com/price-calculator/filemanager"
	"example.com/price-calculator/prices"
)

func main() {
	taxRate := []float64{0, 0.8, 0.6, 0.9, 0.15}

	for _, taxRate := range taxRate { 
		fm := filemanager.New("price.txt", fmt.Sprintf("result_%0f.json", taxRate*100))
		priceJob := prices.NewTaxIncludedPricesJob(fm ,taxRate)
		priceJob.Calculate()
	}
	
}
