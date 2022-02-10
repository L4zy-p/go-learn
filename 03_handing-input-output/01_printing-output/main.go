package main

import "fmt"

func main() {
	name := "L4zy-p"
	var initial rune = 'P'
	age := 18
	height := 160
	hoursWatched := 1.5
	fmt.Printf("Hello %c and welecome %v\n Your age is %d.\n Your height is %d.\n Good for you!\n", initial, name, age, height)
	// %c คือ rune
	// %v จะเป็นพวก default format
	// %d จะเป็นพวกตัวเลข
	result := fmt.Sprintf("You've watched %10.2f hours form video", hoursWatched)

	fmt.Println("Lecture breakdown:")
	fmt.Printf("Introduction to programing | %10.2f |\n", 0.67)
	fmt.Printf("Introduction to programing | %5.2f |\n", 0.67)
	fmt.Printf("Introduction to programing | %.3f |\n", 0.67)
	// %f จะเป็นพวกทศนิยม โดยที่
	// หลัง . จะเป็นกำหนดตำแหน่งทศนิยม
	// หน้า . หลัง % จะเป็นการกำหนด space หน้า

	fmt.Println(result)
}
