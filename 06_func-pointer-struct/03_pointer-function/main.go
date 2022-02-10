package main

import (
	"fmt"
)

func main() {
	start, end := 10, 20
	iterate(start, end)
	fmt.Println(start, end) // 10 20

	iteratePtr(&start, &end)
	fmt.Println(start, end) // 20 20

	// จะเห็นว่าตัวที่ใช้ pointer func จะเปลี่ยนค่าที่เราชี้ไปด้วย
}

func iterate(start, end int) {
	for start < end {
		fmt.Println(start)
		start++
	}
}

func iteratePtr(start, end *int) {
	for *start < *end {
		fmt.Println(*start)
		(*start)++
	}
}
