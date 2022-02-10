package main

import "fmt"

func main() {
	// บวก ลบ คูณ หาร mod
	num1 := 14
	num2 := 24

	sum := num1 + num2
	subtr := num2 - num1
	multi := num1 * num2
	divide := num2 / num1
	modulus := num2 % num1

	fmt.Println("sum", sum)         // sum 38
	fmt.Println("subtr", subtr)     // subtr 10
	fmt.Println("multi", multi)     // multi 336
	fmt.Println("divide", divide)   // divide 1
	fmt.Println("modulus", modulus) // modulus 0

	num1++ // num1 15
	num2-- // num2
	fmt.Println("num1", num1)
	fmt.Println("num2", num2)

	// short operator
	var a, b, c, d int
	a += 5 // a = a + 5
	b -= 3 // b = b - 3
	c *= 2 // c = c * 2
	d /= 2 // d = d / 2
}
