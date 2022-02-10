package main

import (
	"fmt"
	"strings"
)

func main() {
	// operator ในการเปรียบเทียบ
	num := 15
	answer := num == 12
	fmt.Println("Answer is", answer)

	fmt.Println("Please enter you name")
	var name string
	fmt.Scan(&name)
	correctName := strings.TrimSpace(name) == "Lazy"
	fmt.Println("You are name Lazy", correctName)

	message := "Hello there Mr. Smith"
	// contains เป็น func ที่แยกออกมาเป็นคำๆ เพื่อเช็ค ถ้าตำแหน่งไหนคำถูกหรือว่าง ก็จะเป็น true
	checkMessageContain := strings.Contains(message, "Hello there Ms.")
	fmt.Println("Is contains message", checkMessageContain)
}
