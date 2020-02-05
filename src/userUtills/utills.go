package userutills

import "fmt"

//utility functions

//GetStringInput returns a string and takes a message to display
func GetStringInput(msg string) string {
	fmt.Print(msg)
	var input string
	fmt.Scanln(&input)

	return input
}

//GetIntInput returns an int and takes a message to display
func GetIntInput(msg string) int {
	fmt.Print(msg)
	var input int
	fmt.Scanln(&input)

	return input
}
