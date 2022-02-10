package main

import (
	"fmt"
	"math"
)

func main() {
	// math เป็น Lib ที่เกี่ยวกับการคำนวนจะมี func ที่ช่วยในการคำนวน
	var num int = 9
	result := math.Sqrt(float64(num)) // func ส่วนใหญ่ใน Math จะรับค่าเป็น type float64 จึงทำให้ตรงแปลง type ก่อนส่ง
	fmt.Println(result)               // 3

	num = -14
	fmt.Println(math.Abs(float64(num))) // 14

	num = 9
	fmt.Println(math.Pow(float64(num), 3)) // 729
}
