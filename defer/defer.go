package main

import "fmt"

func add(value1 int, value2 int) int {
	return value1 + value2
}

func loop() {
	for i := 0; i < 10; i++ {
		fmt.Println("i =", i)
	}
}

func deferLoop() {
	for j := 0; j < 10; j++ {
		defer fmt.Println("j =", j)
	}
}

func main() {
	// fmt.Println("Welcome to calculator.")
	// defer fmt.Println("End")
	// defer fmt.Println(add(20, 10))
	// defer fmt.Println(add(15, 15))
	// defer fmt.Println(add(12, 12))

	loop()
	deferLoop()
}

/*
Summary

`defer` เป็นคีย์เวิร์ดที่ใช้ "เลื่อน" การทำงานของฟังก์ชันไว้ทำทีหลัง
โดยฟังก์ชันที่ถูก `defer` จะถูกเรียกทำงาน **ก่อนที่ฟังก์ชันที่ครอบมันอยู่จะจบการทำงาน (return)**

หลักการทำงานสำคัญ:

1.  การทำงานแบบ LIFO (Last-In, First-Out):
    - การเรียก `defer` หลายๆ ครั้ง จะเหมือนการนำฟังก์ชันไปเก็บไว้ในกองซ้อน (Stack)
    - เมื่อฟังก์ชันหลักกำลังจะจบ, Go จะดึงฟังก์ชันที่ถูก `defer` ไว้จากกองซ้อนออกมาทำงานทีละตัว **จากตัวล่าสุดไปหาตัวแรกสุด**
    - ในฟังก์ชัน `deferLoop()`: `defer fmt.Println("j =", 9)` ถูกเรียกเป็นตัวสุดท้าย แต่มันจะทำงานเป็นตัวแรก
      ทำให้ผลลัพธ์ที่แสดงออกมาเป็น `j = 9` ไล่ลงไปจนถึง `j = 0`

2.  การประเมินค่า Argument ทันที (Immediate Argument Evaluation):
    - Argument ของฟังก์ชันที่ถูก `defer` (เช่น ค่า `j` ใน `fmt.Println("j =", j)`) จะถูกประเมินค่าและเก็บไว้ **ณ เวลาที่เรียก `defer`**
      ไม่ใช่ตอนที่ฟังก์ชันถูกเรียกทำงานจริง นี่คือเหตุผลที่มันสามารถพิมพ์ค่า `j` ที่ถูกต้องของแต่ละรอบได้

3.  ประโยชน์การใช้งาน:
    - `defer` มีประโยชน์มากในการจัดการทรัพยากร เช่น การปิดไฟล์ (`file.Close()`), การปลดล็อก mutex (`mutex.Unlock()`)
      การใช้ `defer` ช่วยให้โค้ดสะอาดขึ้นและลดโอกาสลืมคืนทรัพยากร เพราะเราสามารถเขียนโค้ดเปิดทรัพยากรและ `defer` การปิดไว้คู่กันได้เลย
*/
