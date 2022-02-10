package main

import "fmt"

func main() {
	// int จะเป็นเลขที่ไม่เกิน 19 หลัก
	var num int = -999999999999999999
	fmt.Println(num)

	// int8 จะเป็นช่วง -128 ถีง 127
	var num2 int8 = 127
	fmt.Println(num2)

	// int16 จะเป็นช่วง -32768 ถีง 32767
	var num3 int16 = 32767
	fmt.Println(num3)

	// int32 จะเป็นช่วง -2147483648 ถีง 2147483647
	var num4 int32 = 2147483647
	fmt.Println(num4)

	// int64 จะเป็นช่วง -9223372036854775808 ถีง 9223372036854775807
	var num5 int64 = 9223372036854775807
	fmt.Println(num5)

	// พวกที่เป็น uint จะไม่สามารถเป็นลบได้
	var notnegative uint = 9999999999999999999
	fmt.Println(notnegative)

	// uint8 จะเป็นช่วง 0 ถีง 255 นอกจากนี้ type นี้สามารถเป็น byte ได้
	var notnegative2 uint8 = 255
	// var notnegative2 byte = 255
	fmt.Println(notnegative2)

	// uint16 จะเป็นช่วง 0 ถีง 65535
	var notnegative3 uint16 = 65535
	fmt.Println(notnegative3)

	// uint32 จะเป็นช่วง 0 ถีง 4294967295
	var notnegative4 uint32 = 4294967295
	fmt.Println(notnegative4)

	// uint32 จะเป็นช่วง 0 ถีง 18446744073709551615
	var notnegative5 uint64 = 18446744073709551615
	fmt.Println(notnegative5)
}
