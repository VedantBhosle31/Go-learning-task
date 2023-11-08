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
	fmt.Scan(&user.FirstName)
	fmt.Printf("Enter your Last Name ")
	fmt.Scan(&user.LastName)
	fmt.Printf("Enter your email ")
	fmt.Scan(&user.Email)
	fmt.Printf("Enter how mant tickets your want ")
	fmt.Scan(&user.RequestedTickets)
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
