package main

import "fmt"

func main() {
	var aFloat float32 = 14.3
	fmt.Println("sum", aFloat+3)    // sum 17.3
	fmt.Println("subtr", aFloat-3)  // subtr 11.3
	fmt.Println("multi", aFloat*3)  // multi 42.9
	fmt.Println("divide", aFloat/3) // divide 4.766667 // จำนวนทศนิยมที่แสดงจะเท่ากับจำนวน bit ที่แสดงได้

	// mod จะไม่ได้กำหนดมาเพื่อให้ใช้กับ float

	var bFloat float32 = aFloat
	aFloat++
	bFloat--
	fmt.Println("aFloat++", aFloat) // aFloat++ 15.3
	fmt.Println("bFloat--", bFloat) // bFloat-- 13.3

	var num1 float64 = 3.0
	num1 = num1 + 0.1
	num1 = num1 + 0.1
	num1 = num1 + 0.1
	num1 = num1 + 0.1
	num1 = num1 + 0.1
	num1 = num1 + 0.1
	num1 = num1 + 0.1
	num1 = num1 + 0.1
	fmt.Println("num1", num1) // num1 3.8000000000000007

	var divideZero = 3.0
	fmt.Println(divideZero / 0.0)   // +Inf
	fmt.Println(divideZero / -0.0)  // +Inf
	fmt.Println(-divideZero / 0.0)  // -Inf
	fmt.Println(-divideZero / -0.0) // -Inf
	// จากเรื่อง integer operation จะหารด้วย 0 ไม่ได้ แต่ float จะทำได้โดยไม่แจ้ง error
	// เมื่อหารด้วย 0.0 แล้วจะได้ค่าเป็น inf ทุกตัว
	// จะเป็น + หรือ - ขึ้นอยู่กับตัวทั้งว่าเป็น + หรือ -
}
