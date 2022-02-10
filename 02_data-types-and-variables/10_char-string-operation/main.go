package main

import "fmt"

func main() {
	var ch rune = '5'
	fmt.Println(ch) // ที่เป็น 53 คือเลข ascii

	// ดังนั้นเมื่อเป็นตัวเลขจึงสามารถใช้ operator ได้
	ch = ch + 1
	fmt.Println(ch)     // 54
	fmt.Println(ch / 2) // 27
	fmt.Println(ch * 2) // 108
	fmt.Println(ch % 2) // 0

	var message string = "Go is "
	message = message + "awesome!!!!"
	fmt.Println(message) // Go is awesome!!!!
	// การบวกใน string จะเป็นการเอามาต่อกัน

	// เพิ่มเติมเรื่อง escaping character

	fmt.Println("Go is\n awesome")            // \n เป็นการขึ้นบรรทัดใหม่
	fmt.Println("Go is\t awesome")            // \t เป็นการ tab
	fmt.Println("Go isawesome\n \"cool!!!\"") // หากต้องการใส้ " ให้ใส่ \"
}
