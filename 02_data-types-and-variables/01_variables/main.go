package main

import "fmt"

func main() {
	// ใน go การกำหนด variable const var ขึ้นมาสักตัวจะต้องกำหนด type ให้ด้วยถ้าไม่ได้กำหนดค่าเรื่มต้น
	var message string
	message = "Hello world"

	fmt.Println(message)

	// message2 ไม่กำหนด Type ก็ได้เพราะ Go จะกำหนด type ให้เองจากค่าเริ่มต้น
	var message2 = "Golang"

	fmt.Println(message2)

	// message3 จะเหมือน message2 แต่เขียนในรูปแบบสั้น
	message3 := "learing"

	fmt.Print(message3)
}
