package main

import (
	"booking-cli-in-go/booking-cli-in-go/helper"
	"fmt"
	"strings"
)

type UserData helper.UserData

func main() {

	for {

		fmt.Println("Welcome to the Go Confrence")
		const totalTickets int = 50
		var remainingTickets uint = 50
		fmt.Printf("We have only %v tickets left.\nPlease Register to Confirm your booking", remainingTickets)
		// firstName, lastName, email, requestedTickets
		// var user UserData
		user := helper.GetInput()

		isValidEmail := strings.Contains(user.Email, "@")

		isValidNumber := user.RequestedTickets > remainingTickets

		//validation for ticket sales
		if isValidEmail && isValidNumber {
			if isValidEmail {
				fmt.Println("Invalid Email")
			}
			if isValidNumber {
				fmt.Printf("%v tickets not available\n", user.RequestedTickets)
			}
			continue
		} else {

			remainingTickets = remainingTickets - user.RequestedTickets

			fmt.Printf("Thank you %v %v for booking % v Tickets.\n", user.FirstName, user.LastName, remainingTickets)
			fmt.Printf("You will recevie a confirmation mail on %v \n", user.Email)

			var allTheAttendes []string
			allTheAttendes = append(allTheAttendes, user.FirstName+" "+user.LastName)
			firstNames := helper.GetFirstName(allTheAttendes)

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
