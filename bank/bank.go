package main

import "fmt"

func main() {
	var balanceAmount float64 = 1000

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit Money")
	fmt.Println("3. Withdraw Money")
	fmt.Println("4. Exit")

	var choice int

	fmt.Print("Enter your choice:  ")
	fmt.Scan(&choice)

	if choice == 1 {
		fmt.Println("Your balance is :", balanceAmount)
	} else if choice == 2 {
		fmt.Print("Enter the deposit amount: ")
		var depositAmount float64
		fmt.Scan(&depositAmount)
		balanceAmount += depositAmount
		fmt.Println("Your new balance is :", balanceAmount)
	} else if choice == 3 {
		fmt.Print("Enter the withdrawal amount: ")
		var withdrawalAmount float64
		fmt.Scan(&withdrawalAmount)
		balanceAmount -= withdrawalAmount
		fmt.Println("Your new balance is :", balanceAmount)
	} else {
		fmt.Println("Goodbye!")
	}
}
