package main

import "fmt"

func main() {
	input := 4

	switch input {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Unknown number")
	}
}
