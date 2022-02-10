package main

import "fmt"

var number = 14

func main() {
	foo(2)
	fmt.Println(number)
}

func foo(a int) {
	number = 20
	fmt.Println(number)
}
