package main

import (
	"fmt"
)

func main() {
	var birthYear int
	var currentYear int

	fmt.Println("Please, enter your birth year: ")
	_, err := fmt.Scanln(&birthYear)

	if err != nil {
		fmt.Println("Invalid year has been input. Please enter a valid year.")
		return
	}

	fmt.Println("Please, enter your current year: ")
	_, err = fmt.Scanln(&currentYear)
	if err != nil {
		fmt.Println("Invalid current year has been input. Please enter a valid current year.")
		return
	}

	if birthYear > currentYear {
		fmt.Println("Birth year cannot be greater than the current year.")
		return
	}

	age := currentYear - birthYear

	fmt.Printf("Your age is: %d\n", age)
}
