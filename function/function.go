package main

func hello() {
	println("Hello Diego")
}

func plus(a int, b int) int {
	return a + b
}

func plus3value(a, b, c int) int {
	return a + b + c
}

func main() {
	hello()
	result := plus(10, 20)
	println(result)

	result = plus3value(10, 20, 30)
	println(result)
}
