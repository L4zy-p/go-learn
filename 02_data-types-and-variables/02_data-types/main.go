package main

import (
	"fmt"
	"time"
)

func main() {
	// text (ข้อความ)
	var message string = "Hello go lang"
	fmt.Println(message)

	// whole number (จำนวนเต็ม)
	var number int = 13
	fmt.Println(number)

	// real number (ทศนิยม)
	var realNumber float32 = 1.99
	fmt.Println(realNumber)

	// character (ตัวอักษร)
	var character rune = 'P'
	fmt.Println(character)

	// boolean
	var isTrue bool = true
	fmt.Println(isTrue)

	// สามารถเรียก type อื่นๆ มากำหนด type ได้
	var currentTime time.Time = time.Now()
	fmt.Println(currentTime)
}
