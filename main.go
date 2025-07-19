package main

import "fmt"

var numberInt, numberInt2 int = 1000, 2000
var msg string = "Hello World"

func main() {
	numberFloat := 25.4
	fmt.Println(numberInt)
	fmt.Println(numberInt2)
	fmt.Println(numberFloat)
	fmt.Println(msg)

	fmt.Println(float64(numberInt) + numberFloat)
	fmt.Println(msg + " My name is Diego")
	fmt.Println("my money =", numberInt)
}
