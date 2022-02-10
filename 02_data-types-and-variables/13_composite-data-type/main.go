package main

import (
	"fmt"
	"time"
)

func main() {
	// composite data type คือข้อมูลเชิงประกอบ สามารถเก็บข้อมูลได้หลายตัวแปรในตัวเดียวกัน
	var t1 time.Time
	t1 = time.Now()

	year, month, day := t1.Date()
	fmt.Println(year, month, day)
}
