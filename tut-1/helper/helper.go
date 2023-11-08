package helper

import "fmt"

func GetInput() (string, string, string, uint) {
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
