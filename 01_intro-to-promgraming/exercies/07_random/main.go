package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// ให้ rand เรียก seed เพื่อใช้ nanosecond ปัจจุบัน
	rand.Seed(int64(time.Now().Nanosecond()))

	// random int จะต่างกันในแต่ละเวลาที่ execute จะ random ตัวเลขที่ไม่เกิน 11
	random := rand.Intn(11)
	fmt.Println(random)
}
