package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	investmentAmount, expectedReturnRate := 0.0, 5.5
	var years float64

	outputText("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	outputText("Years: ")
	fmt.Scan(&years)

	outputText("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue, futureRealValue := calculateFutureValue(investmentAmount, expectedReturnRate, years)

	formattedFv := fmt.Sprintf("Future Value: %.0f", futureValue)
	formattedFrv := fmt.Sprintf("Future Real Value: %.0f\n", futureRealValue)

	fmt.Print(formattedFv, formattedFrv)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValue(investmentAmount, expectedReturnRate, years float64) (float64, float64) {
	fv := investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	frv := fv / math.Pow((1+inflationRate/100), years)

	return fv, frv
}
