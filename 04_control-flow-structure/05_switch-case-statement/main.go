package main

import (
	"fmt"
	"strings"
)

func main() {
	var dayOfWeek string
	fmt.Println("Please enter day of week")
	fmt.Scan(&dayOfWeek)

	dayOfWeek = strings.TrimSpace(dayOfWeek)

	switch dayOfWeek {
	case "monday":
		// fmt.Println("It's monday")
		fallthrough // ถ้าเป็น monday ให้ทำใน case
	case "tuesday":
		fmt.Println("It's tuesday")
	case "wednesday":
		fmt.Println("It's wednesday")
	case "thursday":
		fmt.Println("It's thursday")
	case "friday":
		fmt.Println("It's friday")
	case "saturday":
		fmt.Println("It's saturday")
	case "sunday":
		fmt.Println("It's sunday")
	default:
		fmt.Println("It's weekend")
	}

	// ถ้าเราอยากให้ fallthrough สามารถเขียนแบบด้านล่างได้อีกรูปแบบ
	// switch dayOfWeek {
	// case "monday":
	// case "tuesday":
	// case "wednesday":
	// 	fmt.Println("It's wednesday")
	// }
}
