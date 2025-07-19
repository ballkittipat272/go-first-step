package main

import (
	"fmt"
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
}

/*
Summary

Channel คือ "ท่อ" สื่อสารที่ปลอดภัยสำหรับให้ Goroutine ใช้ส่งและรับข้อมูลระหว่างกัน

หลักการทำงาน:
1.  การสร้าง: `ch := make(chan string)` สร้าง channel สำหรับส่งข้อมูลประเภท string
    - Channel ที่สร้างแบบนี้เรียกว่า "Unbuffered Channel"

2.  การส่งและรับ:
    - ส่งข้อมูล: `ch <- data` (ส่งค่า `data` เข้าไปใน channel)
    - รับข้อมูล: `msg := <-ch` (รอรับค่าจาก channel มาเก็บใน `msg`)

3.  การทำงานแบบ Blocking และ Synchronization (หัวใจสำคัญ):
    - การส่ง (send) และการรับ (receive) บน Unbuffered Channel จะ "หยุดรอ" (block)
    - ผู้ส่งจะรอจนกว่าจะมีผู้รับมารับ และผู้รับจะรอจนกว่าจะมีผู้ส่งมาส่ง
    - คุณสมบัตินี้ทำให้ Channel เป็นเครื่องมือที่ยอดเยี่ยมในการจัดลำดับการทำงาน (Synchronize) ของ Goroutine
    - ในโค้ดนี้ `main` จะหยุดรอที่ `<-ch` จนกว่า Goroutine `process1` จะส่งข้อมูลมาให้ ทำให้เรามั่นใจได้ว่า `process1` ได้ทำงานแล้ว โดยไม่จำเป็นต้องใช้ `time.Sleep`
*/
