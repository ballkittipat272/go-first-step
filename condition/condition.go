package main

import "fmt"

func calculategrade(score int) string {
	if score >= 80 {
		return "A"
	} else if score >= 70 {
		return "B"
	} else if score >= 60 {
		return "C"
	} else if score >= 50 {
		return "D"
	} else {
		return "F"
	}
}

func main() {
	var score int
	fmt.Println("grade calculator")
	fmt.Scanf("%d", &score)
	fmt.Println("score =", score)

	fmt.Println(calculategrade(score))
}
