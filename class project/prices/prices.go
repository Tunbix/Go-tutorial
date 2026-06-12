package prices

type TaxIncludedPricesJob struct {
	TaxRates []float64
	InputPrices   []float64
	TaxIncludedPrices map[string][]float64
}

func (job *TaxIncludedPricesJob) Calculate() {
	

	for _, taxRate := range taxRate {
		taxIncludedPrice := make([]float64, len(prices))
		for priceIndex, price := range prices {
			taxIncludedPrice[priceIndex] = price * (1 + taxRate)
		}
		result[taxRate] = taxIncludedPrice
	 }
}

func NewTaxIncludedPricesJob(taxRate float64) *TaxIncludedPricesJob {
	return &TaxIncludedPricesJob{
		InputPrices:   []float64{10.5, 20.0, 15.75, 30.0, 25.5},
		TaxRates: taxRate,
}