package main

import (
	"fmt"
	"math"
)

func main() {
	 investmentAmount  := 1000.0
	 expectedReturnRate := 5.5
	 Years  := 10.0

	 futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, Years)
	 futureRealValue := futureValue / 
	fmt.Println(futureValue)
}
