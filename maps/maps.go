package main

import "fmt"

var product = make(map[string]float64)

func main() {
	fmt.Println("product =", product)

	// add
	product["Macbook"] = 40000
	product["iPad"] = 30000
	fmt.Println("product =", product)

	// delete
	delete(product, "iPad")
	fmt.Println("product =", product)

	// update
	product["Macbook"] = 50000
	fmt.Println("product =", product)

	value1 := product["Macbook"]
	fmt.Println("value1 =", value1)

	courseName := map[string]string{"101": "Golang", "102": "Python", "103": "Java"}
	fmt.Println("courseName =", courseName)
}
