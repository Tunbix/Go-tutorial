package main

import "fmt"

func main() {
	var accountBalance = 5000.0


	fmt.Println("Welcome to Chase bank")
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit Money")
	fmt.Println("3. Withdraw Money")
	fmt.Println("4. Exit")

	var choice int
	fmt.Println("Your choice:", choice)
	fmt.Scan(&choice)

	//wantsCheckBalance := choice == 1

	if choice == 1 {
		fmt.Println("Your balance is", accountBalance)
	} else if choice == 2 {
		fmt.Print("Your deposit: ")
		var depositAmount float64
		fmt.Scan(&depositAmount)

		
		accountBalance += depositAmount // or use this accountBalance = accountBalance + depositAmount
		fmt.Println("Balance updated! New amount:", accountBalance)
	} else if choice == 3 {
		fmt.Print("Withdrawal amount: ")
		var WithdrawalAmount float64
		fmt.Scan(&WithdrawalAmount)
		accountBalance -= WithdrawalAmount
		fmt.Println("Your current accout balance:", accountBalance)
	} else {
		fmt.Println("Thank you!")
	}


	fmt.Println("Your choice:", choice)

	

	

}
