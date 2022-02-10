package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("numbers", numbers)
	subSlice := numbers[:2]       // จาก index ที่ 0 และนับ index 0 ไปตำแหน่งที่สอง คือ 1 , 2
	subSliceEnd := numbers[2:]    // จาก index ที่ 2 และนับ index 2 ไปตัวสุดท้าย ตัว คือ 3,4,5,6
	subSliceMidle := numbers[2:4] // จาก index ที่ 0 และนับ index 0 ไปตำแหน่งที่ 4 คือ 3, 4
	fmt.Println("sub slice 0,2", subSlice)
	fmt.Println("sub slice 0,last", subSliceEnd)
	fmt.Println("sub slice 2,4", subSliceMidle)
}
