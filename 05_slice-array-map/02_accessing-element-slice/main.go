package main

import "fmt"

func main() {
	zombies := []string{"Paul", "Gooey", "George", "Lucy"}
	paul := zombies[0]
	index := 2
	george := zombies[index]

	fmt.Println(zombies, paul, george)

	zombies[1] = "Katya"

	fmt.Println("zombies 1", zombies[1])
	fmt.Println("all zombies", zombies, "have", len(zombies), "zombies")
}
