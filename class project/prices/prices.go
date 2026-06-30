package prices

import (
	"fmt"
	
	"example.com/price-calculator/filemanager"
	"example.com/price-calculator/conversion"
)

type TaxIncludedPricesJob struct {
	IOManager		filemanager.FileManager
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]string
}

func (job *TaxIncludedPricesJob) LoadData(){

	lines, err := job.IOManager.Realines()

	if err != nil {
		fmt.Println(err)
		return
	}

	prices, err := conversion.StringToFloats(lines)

		if err != nil {
			fmt.Println(err)
			
			return
		}
	
	job.InputPrices = prices
	
}


func (job *TaxIncludedPricesJob) Calculate() {
	job.LoadData()
	
	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	job.TaxIncludedPrices = result

	job.IOManager.WriteResult(job)
}

func NewTaxIncludedPricesJob(fm filemanager.FileManager, taxRate float64) *TaxIncludedPricesJob {
	return &TaxIncludedPricesJob{
		InputPrices: []float64{10.5, 20.0, 15.75, 30.0, 25.5},
		TaxRate:     taxRate,
	}
}
