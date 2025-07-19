package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 100; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	go f("Hello from BALL...")
	go f("Work number 2")
	time.Sleep(5 * time.Second)
}

/*
Summary

Goroutine คือหน่วยการทำงานย่อย (lightweight thread) ที่ถูกจัดการโดย Go runtime
ทำให้เราสามารถรันฟังก์ชันหลายๆ ตัวพร้อมกันได้ (Concurrency)

วิธีการใช้งาน:
- สร้าง Goroutine ได้ง่ายๆ ด้วยการใช้คีย์เวิร์ด `go` นำหน้าฟังก์ชันที่ต้องการรัน
  เช่น `go f("Hello from BALL...")` คือการสั่งให้ฟังก์ชัน `f` ทำงานใน Goroutine ใหม่

การทำงานพร้อมกัน (Concurrency):
- จากโค้ด `main` จะสร้าง Goroutine ใหม่ 2 ตัว (`go f(...)`) ซึ่งจะทำงานไปพร้อมๆ กับ Goroutine หลัก (main)
- ผลลัพธ์ที่พิมพ์ออกมาจะสลับกันไปมา เพราะทั้งสอง Goroutine แข่งกันทำงาน

ความสำคัญของ `time.Sleep`:
- เมื่อ Goroutine หลัก (main) ทำงานจบ, โปรแกรมทั้งหมดจะปิดตัวลงทันที
- `time.Sleep` ในที่นี้ทำหน้าที่ "ถ่วงเวลา" เพื่อให้ Goroutine อื่นๆ มีเวลาทำงานจนเสร็จ นำมาช่วยเป็นตัวแสดงผลตัวอย่าง
- ในแอปพลิเคชันจริง ควรใช้ Channels หรือ WaitGroups เพื่อรอการทำงานของ Goroutine จะดีกว่า
*/
