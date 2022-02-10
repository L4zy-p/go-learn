package main

import "fmt"

func main() {
	// bool จะเป็น type true หรือ false
	// rune เป็นพวกตัวอักษรหรือตัวเลขที่อยู่ในรูปแบบ string แค่ตัวเดียวเท่านั้น ลองใส่ อิโมจิแล้วออกเป็นตัวเลข
	// string จะเป็พวก text จะใส่พวกอิโมจิได้
	var isTrue bool = false
	var ch rune = '😀'
	var text string = "Hello world 😀"

	fmt.Print(isTrue, ch, text) // false 128512Hello world 😀
}
