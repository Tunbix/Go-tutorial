package main

import(
	"fmt"
	"time"
)

type user struct{
	firstName string
	lastName string
	birthdate string
	createdAt time.Time
}

func main() {

	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdaten (MM/DD/YYYY): ")

	var appUser user

	appUser = user{
		firstName: userFirstName,
		lastName: userLastName,
		birthdate: userBirthdate,
		createdAt: time.Now(),
	}

	fmt.Println(firstName, lastName, birthdate)

}

func getUserData(promptText string) string {

	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
	
}