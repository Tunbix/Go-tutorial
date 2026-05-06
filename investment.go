package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	 var  investmentAmount float64
	 expectedReturnRate := 5.5
	 var Years  float64

	 fmt.Print("Investment Amount: ")
	 fmt.Scan(&investmentAmount)

	 fmt.Print("Return Rate: ")
	 fmt.Scan(&expectedReturnRate)

	 fmt.Print("Years: ")
	 fmt.Scan(&Years)


	 futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, Years)
	 futureRealValue := futureValue /  math.Pow(1+inflationRate/100, Years)

	 // output information 
	 fmt.Println("Future Value:", futureValue)
	 fmt.Println(futureRealValue)
}
