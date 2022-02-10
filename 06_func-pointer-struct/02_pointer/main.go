package main

import "fmt"

func main() {
	num := 10
	fmt.Println(num)

	numCopy := num
	fmt.Println(numCopy)

	numCopy = 15
	fmt.Println(num, numCopy)

	var numPointer *int           // สร้าง Pointer int
	numPointer = &num             // ให้ pointer ชี้ไปที่ num
	fmt.Println(num, *numPointer) // ในการกึงค่าของ pointer ต้องใส่ * ถ้าไม่ใส่จะได้เลข address แทน
	*numPointer = 20              // เมื่อ Pointer เปลี่ยนค่า num ก็จะเปลี่ยนด้วย
	fmt.Println(num, *numPointer)
}
