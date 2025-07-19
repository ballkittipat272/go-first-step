package main

import "fmt"

func main() {
	var colorText string
	fmt.Print("Enter color: ")
	fmt.Scanln(&colorText)

	switch colorText {
	case "blue":
		fmt.Println("#0000FF")
	case "green":
		fmt.Println("#008000")
	case "pink":
		fmt.Println("#FFC0CB")
	case "yellow":
		fmt.Println("#FFFF00")
	default:
		fmt.Println("Unknown color")
	}
}
