package main

import (
	"fmt"
	"strings"
)

func main() {

	for {

		fmt.Println("Welcome to the Go Confrence")
		const totalTickets int = 50
		var remainingTickets uint = 50
		fmt.Printf("We have only %v tickets left.\nPlease Register to Confirm your booking", remainingTickets)
		firstName, lastName, email, requestedTickets := getInput()
		isValidEmail := strings.Contains(email, "@")

		isValidNumber := requestedTickets > remainingTickets

		//validation for ticket sales
		if isValidEmail && isValidNumber {
			if isValidEmail {
				fmt.Println("Invalid Email")
			}
			if isValidNumber {
				fmt.Printf("%v tickets not available\n", requestedTickets)
			}
			continue
		} else {

			remainingTickets = remainingTickets - requestedTickets

			fmt.Printf("Thank you %v %v for booking % v Tickets./n", firstName, lastName, remainingTickets)
			fmt.Printf("You will recevie a confirmation mail on %v \n", email)

			var allTheAttendes []string
			allTheAttendes = append(allTheAttendes, firstName+" "+lastName)
			firstNames := getFirstName(allTheAttendes)

			fmt.Printf("The List of all the attendees is %v.\n", firstNames)
			fmt.Printf("Number people who booked the tickets is%v.\n", len(firstNames))
		}
		if remainingTickets == 0 {
			fmt.Println("We are sold out")
			break
		} else {
			fmt.Printf("Number of remaining Tickets is %v\n", remainingTickets)
		}
	}
}

func getFirstName(allTheAttendes []string) []string {
	firstNames := []string{}

	for _, attendee := range allTheAttendes {
		var name = strings.Fields(attendee)
		firstNames = append(firstNames, name[0])
	}
	return firstNames
}

func getInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var requestedTickets uint

	fmt.Printf("Enter your First Name ")
	fmt.Scan(&firstName)
	fmt.Printf("Enter your Last Name ")
	fmt.Scan(&lastName)
	fmt.Printf("Enter your email ")
	fmt.Scan(&email)
	fmt.Printf("Enter how mant tickets your want ")
	fmt.Scan(&requestedTickets)
	return firstName, lastName, email, requestedTickets
}
