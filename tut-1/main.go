package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")
	// Phase !
	var ConferenceName = "Go confrence"
	fmt.Printf("Welcome to %v \n", ConferenceName)

	//Phase 2

	var userName string
	userName = "Vedant"
	fmt.Println(userName)

	//Phase 3 Getting types

	fmt.Printf(" The type of the variable %v is %T \n", userName, userName)

	//Phase 4
	var userid string
	fmt.Println("Enter your userid")
	fmt.Scan(&userid)

	fmt.Printf("We user has been set to %v \n", userid)
}
