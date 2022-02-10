package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// bufio.NewReader จะอนุญาตให้อ่าน text line input
	fmt.Println("Please enter number")
	r := bufio.NewReader(os.Stdin)
	line, err := r.ReadString('\n') // จะอ่าน text line เมื่อมีการขึ้น line ใหม่
	if err != nil {
		panic(err)
	}

	// แปลง text ที่ได้รับมาให้เป็น Number และตัด space ออก
	num, err := strconv.Atoi(strings.TrimSpace(line))
	if err != nil {
		panic(err)
	}

	fmt.Println("my number is", num)
	fmt.Println("increment by 5", num+5)

	// การรับ input จาก fmt.Scan
	fmt.Println("Please enter num1")
	var num1 int
	fmt.Scan(&num1)
	fmt.Println(num1)

	fmt.Println("Please enter num1")
	var num2 float64
	fmt.Scan(&num2)
	fmt.Println(num2)
}
