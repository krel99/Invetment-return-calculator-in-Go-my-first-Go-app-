package main

import (
	"fmt"
	"math"
)

// go run investment_calculator.go

func main(){
	const inflationRate = 3.0;
	var investmentAmount float64 = 1000; // can also define multiple on a single line
	expectedReturnRate := 4.5;
	var years int32;

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount) // & binds the value to the default

	fmt.Print("Years of deposit: ")
	fmt.Scan(&years)

	fmt.Print("Expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	var futureValue = investmentAmount * math.Pow((1 + expectedReturnRate / 100), float64(years)) 
	var futureRealValue = futureValue / math.Pow(1+inflationRate/100, float64(years))

	formattedFV := fmt.Sprintf("%.1f\n", futureRealValue)

	fmt.Printf("Future value: %.1f\n", futureValue)
	fmt.Println("Future value considering 3.0% inflation", formattedFV)
}

