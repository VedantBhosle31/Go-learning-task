package helper

import (
	"fmt"
	"strings"
	// "tut-1/helper"
)

// type UserData UserData

func GetInput() UserData {
	var user UserData

	fmt.Printf("Enter your First Name ")
	fmt.Scan(&user.firstName)
	fmt.Printf("Enter your Last Name ")
	fmt.Scan(&user.lastName)
	fmt.Printf("Enter your email ")
	fmt.Scan(&user.email)
	fmt.Printf("Enter how mant tickets your want ")
	fmt.Scan(&user.requestedTickets)
	return user
}

func GetFirstName(allTheAttendes []string) []string {
	firstNames := []string{}

	for _, attendee := range allTheAttendes {
		var name = strings.Fields(attendee)
		firstNames = append(firstNames, name[0])
	}
	return firstNames
}
