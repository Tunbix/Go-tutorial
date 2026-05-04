package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	 investmentAmount  := 1000.0
	 expectedReturnRate := 5.5
	 Years  := 10.0

	 futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, Years)
	 futureRealValue := futureValue /  math.Pow(1+inflationRate/100, Years)

	 fmt.Println(futureValue)
	 fmt.Println(futureRealValue)
}
