package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	num := 1
	for num < 10 {
		fmt.Println(num)
		num++
	}

	gameOver := false
	rand.Seed(time.Now().UTC().UnixNano())
	for !gameOver {
		fmt.Println("The game is on-going")
		num := rand.Intn(100)
		if num < 10 {
			gameOver = true
		}
	}
}
