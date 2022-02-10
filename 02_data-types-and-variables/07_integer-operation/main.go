package main

import "fmt"

func main() {
	// overflow ex1
	var smallInt int8 = 100
	smallInt = smallInt + 28
	fmt.Println(smallInt) // -128
	// ที่เป็น -128 เพราะว่า int8 จะอยู่ในช่วง -128 - 127
	// แล้วค่าที่บวกเข้าไปทำให้เป็น 128 แต่ว่าเกินไปแล้ว เลยทำให้ set ค่าจากจุดเริ่่มต้นของค่าบวกก็คือ ค่าลบ
	// ดังนั้น จึงได้ -128

	// overflow ex2
	var smallInt2 int8 = 100
	smallInt2 = smallInt2 + 29
	fmt.Println(smallInt2) // -127
	// ที่เป็น -127 เพราะว่า เกิน -128 เลยเอาที่เกินมา 1 ไปบวกเข้า กับ -128 + 1

	// underflow ex
	var smallInt3 int8 = -127
	smallInt3 = smallInt3 - 2
	fmt.Println(smallInt3) // 127
	// ที่เป็น 127 เพราะว่าเกิน -128 ไปทำให้ต้องกลับจุดเริ่่มต้นของค่าลบก็คือ ค่าบวก
	// ดังนั้น จึงได้ 127

	// negative mod
	var negativeMod int8 = -15
	negativeMod = negativeMod % -2
	fmt.Println(negativeMod)
	// เป็น -1 เพราะว่าตัวตั้งเป็นลบ เงื่อนไขที่ mod จะเป็นลบได้ คือตัวตั้งต้องเป็นลบ ส่วนตัวหารจะเป็นลบหรือบวกก็ได้
}
