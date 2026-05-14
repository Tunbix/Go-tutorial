package main

import (
	"fmt"

	"example.com/bank/fileops"
	"github.com/pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func main() {
	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("-------------")
		//panic("Can't continue, sorry.")
	}

	fmt.Println("Welcome to Chase bank")
	println("Reach us 24/7", randomdata.PhoneNumber())

	for {

		presentOptions()

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

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}

			accountBalance += depositAmount // or use this accountBalance = accountBalance + depositAmount
			fmt.Println("Balance updated! New amount:", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		} else if choice == 3 {
			fmt.Print("Withdrawal amount: ")
			var WithdrawalAmount float64
			fmt.Scan(&WithdrawalAmount)

			if WithdrawalAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}

			if WithdrawalAmount > accountBalance {
				fmt.Println("Insufficient account balance.")
				continue
			}
			accountBalance -= WithdrawalAmount
			fmt.Println("Your current accout balance:", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		} else {
			fmt.Println("Thank you!")
			//return
			break
		}
	}

	fmt.Println("Thank you for choosing our bank")
}
