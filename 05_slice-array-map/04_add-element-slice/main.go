package main

import "fmt"

func main() {
	zombies := []string{}
	fmt.Println("exist zombies", zombies) // []
	zombies = append(zombies, "Paul", "Katya", "Sue")
	zombiesCopy := append(zombies, "Lucy")
	fmt.Println("new zombies", zombies)      // [Paul Katya Sue]
	fmt.Println("copy zombies", zombiesCopy) // [Paul Katya Sue Lucy]

	all := append(zombies, zombiesCopy...)
	fmt.Println("All zombies", all)
}
