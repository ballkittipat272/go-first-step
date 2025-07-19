package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

type triangle struct {
	base, height float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func (t triangle) area() float64 {
	return 0.5 * t.base * t.height
}

func (t triangle) perim() float64 {
	// Assuming it's a right-angled triangle
	return t.base + t.height + math.Sqrt(t.base*t.base+t.height*t.height)
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println("Area = ", g.area())
	fmt.Println("Perimeter = ", g.perim())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}
	t := triangle{base: 6, height: 8}

	measure(r)
	measure(c)
	measure(t)
}

/*
Summary
Interface คือชุดของ "รูปแบบเมธอด" (Method Signatures) ที่กำหนดพฤติกรรมที่คาดหวังจาก Type ต่างๆ พูดง่ายๆ คือมันเป็นเหมือน "สัญญา" ที่บอกว่า "ถ้า Type ไหนอยากจะเป็น Interface ชนิดนี้ จะต้องมีเมธอดเหล่านี้ให้ครบ"
เราสร้าง Interface ขึ้นมาเพื่อกำหนดว่าพฤติกรรมที่จำเป็นมีอะไรบ้าง

go
type geometry interface {
	area() float64
	perim() float64
}
ในที่นี้ geometry คือสัญญาที่บอกว่า Type ใดๆ ก็ตามที่จะเป็น geometry ได้ จะต้องมีเมธอด area() และ perim() ที่คืนค่าเป็น float64

การ Implement แบบไม่ระบุชัดเจน (Implicit Implementation): นี่คือจุดเด่นของ Go ครับ เราไม่จำเป็นต้องใช้คีย์เวิร์ดอย่าง implements เหมือนในภาษาอื่น (เช่น Java, C#) เพียงแค่ struct ของเรามีเมธอดทั้งหมดที่ Interface กำหนดไว้ครบถ้วนและมีรูปแบบตรงกัน Go ก็จะถือว่า struct นั้นได้ implement Interface นั้นโดยอัตโนมัติ

react (ควรเป็น rect) มีเมธอด area() และ perim()
circle มีเมธอด area() และ perim()
triangle มีเมธอด area() และ perim() ดังนั้น ทั้งสาม struct นี้จึงถือว่าเป็น geometry โดยปริยาย
การใช้งานผ่าน Interface (Polymorphism): เราสามารถสร้างฟังก์ชันที่รับพารามิเตอร์เป็น Interface ได้ ทำให้ฟังก์ชันนั้นยืดหยุ่นและทำงานกับ Type ใดๆ ก็ได้ที่ implement Interface นั้นๆ

go
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println("Area = ", g.area())
	fmt.Println("Perimeter = ", g.perim())
}
ฟังก์ชัน measure ไม่สนใจว่า g จะเป็นสี่เหลี่ยม (react), วงกลม (circle), หรือสามเหลี่ยม (triangle) ขอแค่ g มีเมธอด area() และ perim() ตามที่ geometry กำหนดไว้ก็พอ สิ่งนี้ช่วยลดการพึ่งพากันของโค้ด (Decoupling) และทำให้โค้ดนำกลับมาใช้ใหม่ได้ง่ายขึ้น
*/
