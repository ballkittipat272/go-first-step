package main

import "os"

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	data1 := []byte("hello\nBBalLB")
	err := os.WriteFile("file/data.txt", data1, 0644)
	if err != nil {
		panic(err)
	}

	f, err := os.Create("file/employnee_name")
	check(err)
	defer f.Close()

	data2 := []byte("Sira\nManee")
	os.WriteFile("file/employnee_name.txt", data2, 0644)
}
