package main

import "fmt"

func main() {
	// array จะมีการจอง index ไว้
	zombies := [2]string{"Paul", "Sue"}
	fmt.Println(zombies) // [Paul Sue]
	zombies[1] = "Lucy"  // [Paul Lucy]
	fmt.Println(zombies)
}
