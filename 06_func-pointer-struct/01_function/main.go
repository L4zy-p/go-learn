package main

import (
	"errors"
	"fmt"
)

func main() {
	printPlayerMessage()
	printPlayerNameAndAge("Loo", 13)

	printPlayerInfo([]string{"Loo", "Almond"})

	sum := calSumOfDigit(12345)
	fmt.Println("calculate sum of digits", sum)

	// val, err := calSumOfDigitVer2(-1)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(val)
}

// void function
func printPlayerMessage() {
	fmt.Println("player are on the starting line...")
	fmt.Println("the battle is about to begin")
}

func printPlayerNameAndAge(name string, age int) {
	fmt.Println("player is", name)
	fmt.Println("player age is", age)
}

// return function
func calSumOfDigit(num int) int {
	sum := 0
	for num > 0 {
		sum += num % 10
		num /= 10
	}
	return sum
}

// return multi result
func calSumOfDigitVer2(num int) (int, error) {
	if num < 0 {
		return 0, errors.New("negative value are not supported")
	}
	sum := 0
	for num > 0 {
		sum += num % 10
		num /= 10
	}
	return sum, nil
}

func printPlayerInfo(player []string) {
	for _, player := range player {
		fmt.Printf("player %s is cool\n", player)
	}
}
