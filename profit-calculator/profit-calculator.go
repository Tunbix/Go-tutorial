package main

import (
	"fmt"
)

func main() {

	var revenue float64
	var expenses float64
	var taxRate float64

	//fmt.Print("Revenue Income: ")
	outputText("Revenue Income: ")
	fmt.Scan(&revenue)

	//fmt.Print("Expenses: ")
	outputText("Expenses: ")
	fmt.Scan(&expenses)

	//fmt.Print("Tax Rate: ")
	outputText("Tax Rate: ")
	fmt.Scan(&taxRate)

	revenueBeforeTax := revenue - expenses
	revenueAfterTax := revenueBeforeTax * (1 - taxRate/100)
	incomeRatio := revenueBeforeTax / revenueAfterTax
    //outputData information
	//fmt.Println("Tax Ratio:)
	//fmt.Println("Profit:", revenueAfterTax)
	//fmt.Println("Tax Ratio:", incomeRatio)

	//fmt.Printf("RBT: %.0f\nProfit: %.0f\nTax Ratio: %.0f", revenueBeforeTax, revenueAfterTax, incomeRatio) 

	formattedBT := fmt.Sprintf("RBT: %.1f\n", revenueBeforeTax)
	formattedPT := fmt.Sprintf("Profit: %.1f\n", revenueAfterTax)
	formattedTR := fmt.Sprintf("Tax Ratio: %.1f\n", incomeRatio)

	fmt.Print(formattedBT, formattedPT, formattedTR)

	func outputText(text string) {
		fmt.Print(text)

	}
}
