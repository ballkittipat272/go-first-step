package main

import "fmt"

func main() {
	i := 0
	for i < 10 {
		fmt.Println("i =", i)
		i++
	}

	count := 15
	for j := 0; j < count; j++ {
		fmt.Println("j =", j)
	}

	count2 := 0

	for {
		fmt.Println("Hello World", count2)
		if count2 > 10 {
			break
		}
		count2++
	}
}
