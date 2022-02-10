package main

import (
	"fmt"
	"strconv"
)

func main() {
	// การ convers value ไปเป็น type อื่น
	var intNum int8 = 42
	var biggerInt int64 = int64(intNum)
	fmt.Println(biggerInt) // 42

	// var message string = "Hello"
	// var aNum int = int(message) // error
	// fmt.Println(aNum)
	// เราไม่สามารถเปลี่ยน string ไปเป็น ตัวเลขได้

	var bigNum int16 = 300
	var smallNum int8 = int8(bigNum)
	fmt.Println(smallNum) // 44
	// จะเห็นได้ว่าเมื่อเรา convers ค่าใหญ่ไปค่าเล็กจะทำให้ข้อมูลของ data หายไป

	var anInt int16 = 42
	var aFloat float32 = float32(anInt)
	fmt.Println(aFloat) // 42
	// จะเห็นได้ว่าค่ายังคงเหมือนเดิมเมื่อเราเปลี่ยน int ไปเป็น float แสดงว่าข้อมูลไม่มีการหาย

	aFloat = 3.5
	var anotherInt int = int(aFloat)
	fmt.Println(anotherInt) // 3
	// จะเห็นได้ว่าเมื่อเปลี่ยนจาก float ไปเป็น int data บางส่วนจะหายไปก็คือทศนิยม

	var numAsString string = "42"
	// var num int = int(numAsString) // error
	// จะเห็นได้ว่า string number ไม่สามารถแปลงเป็น int โดยใช้ func int ได้

	num, err := strconv.Atoi(numAsString)
	fmt.Println(num, err) // 42 <nil>
	// strconv.Atoi เป็น lib ที่สามรถแปลง string number เป็น number ได้
	// โดยที่จะ return ค่ามาสองตัวคือ number, error

	var numString string = strconv.Itoa(num)
	fmt.Println(numString)
	// strconv.Itoa เป็น lib ที่สามรถแปลง number เป็น string ได้
	// โดยที่จะ return ค่ามาตัวเดียวคือ string
}
