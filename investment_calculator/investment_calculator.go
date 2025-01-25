package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	investmentAmount, expectedReturnRate := 0.0, 5.5
	var years float64

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue := investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	futureRealValue := futureValue / math.Pow((1+inflationRate/100), years)

	formattedFv := fmt.Sprintf("Future Value: %.0f", futureValue)
	formattedFrv := fmt.Sprintf("Future Real Value: %.0f\n", futureRealValue)

	fmt.Print(formattedFv, formattedFrv)
}
