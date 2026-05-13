package main

import (
	"fmt"
	"os"
	"errors"
)

func storeResult(ebt, profit, ratio float64) {
	results := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.3f\n", ebt, profit, ratio )
	os.WriteFile("result.txt", []byte(results), 0644)
}

func main() {
	revenue, err := getUserInput("Revenue: ")

	if err != nil {
		fmt.Println(err)
		return
	}

	expenses, err := getUserInput("Expenses: ")

	if err != nil {
		fmt.Println(err)
		return
	}

	taxRate , err:= getUserInput("Tax Rate: ")
	
	if err != nil {
		fmt.Println(err)
		return
	}

	

	calculateEarings(revenue, expenses, taxRate)
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	
	ratio := ebt / profit
	
	//outputData information
	//fmt.Println("Tax Ratio:)
	//fmt.Println("Profit:", revenueAfterTax)
	//fmt.Println("Tax Ratio:", incomeRatio)

	//fmt.Printf("RBT: %.0f\nProfit: %.0f\nTax Ratio: %.0f", revenueBeforeTax, revenueAfterTax, incomeRatio)

	formattedBT := fmt.Sprintf("RBT: %.1f\n", ebt)
	formattedPT := fmt.Sprintf("Profit: %.1f\n", profit)
	formattedTR := fmt.Sprintf("Tax Ratio: %.1f\n", ratio)

	storeResult(ebt, profit, ratio)

	fmt.Print(formattedBT, formattedPT, formattedTR)

}

func outputText(text string) {
	fmt.Print(text)
}

func calculateEarings(revenue, expenses, taxRate float64) (float64, float64, float64) {
	bt := revenue - expenses
	pt := bt * (1 - taxRate/100)
	tr := bt / pt

	return bt, pt, tr
}

func getUserInput(infoText string) (float64, error)  {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if userInput <= 0{
		return 0, errors.New("Value must be a postive number")
	}

	return userInput, nil
}
