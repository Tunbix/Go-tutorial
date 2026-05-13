package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)

	if err != nil {
		return 1000, errors.New("Failed to find balance file")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 1000, errors.New("Failed to parse stored balance value")
	}

	return balance, nil
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func main() {
	var accountBalance, err = getBalanceFromFile()

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("-------------")
		//panic("Can't continue, sorry.")
	}

	fmt.Println("Welcome to Chase bank")

	for {

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

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}

			accountBalance += depositAmount // or use this accountBalance = accountBalance + depositAmount
			fmt.Println("Balance updated! New amount:", accountBalance)
			writeBalanceToFile(accountBalance)
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
			writeBalanceToFile(accountBalance)
		} else {
			fmt.Println("Thank you!")
			//return
			break
		}
	}

	fmt.Println("Thank you for choosing our bank")
}
