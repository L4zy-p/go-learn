package main

import "fmt"

func main() {
	number := []int{1, 2, 3, 4}

	for i := 0; i < len(number); i++ {
		number[i] += 1
	}
	fmt.Println("print number round 1", number)

	// range จะเหมือน forEach ใน js คือจะ return index กับ element มาให้
	for i, element := range number {
		fmt.Println("element", element)
		number[i] += 1
	}

	fmt.Println("print number round 2", number)
}
