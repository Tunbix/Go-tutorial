package main

import (

	"fmt"

	
)

func main()  {

	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Revenue Income")
	fmt.Scan(&revenue)

	fmt.Print("Expenses")
	fmt.Scan(&expenses)

	fmt.Print("Tax Rate")
	fmt.Scan(&taxRate)
	
	 revenueBeforeTax := revenue - expenses
	 revenueAfterTax := revenueBeforeTax - taxRate
	 incomeRatio := revenueBeforeTax / revenueAfterTax

	 fmt.Println(revenueBeforeTax)
	 fmt.Println(revenueAfterTax)
	 fmt.Println(incomeRatio)

	
}