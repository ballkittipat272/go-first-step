package main

import (
	"fmt"
	"time"
)

func process1(ch chan string, data string) {
	fmt.Println("sending:", data)
	ch <- data
}

func main() {
	fmt.Println("=== main started ===")

	ch := make(chan string)
	go process1(ch, "data1")

	msg := <-ch
	fmt.Println("received from channel:", msg)

	time.Sleep(5 * time.Second)
	fmt.Println("=== main done ===")
}
