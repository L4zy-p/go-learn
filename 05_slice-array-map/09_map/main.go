package main

import (
	"fmt"
)

func main() {
	/************ defining ***********/

	playerHPs := map[string]int{
		"cool": 10,
		"test": 50,
		"www":  30,
	}

	/************ accessing element ***********/

	fmt.Println("=============================================")
	// fmt.Println(playerHPs)
	hp := playerHPs["www"]
	fmt.Println(hp) // 30

	playerHPs["test"] = 25
	fmt.Println(playerHPs["test"]) // 25

	hp, ok := playerHPs["nice"]
	fmt.Println(hp, ok) // 0 false
	if !ok {
		fmt.Println("this player doesn't exist")
	}

	/************ add new element ***********/

	fmt.Println("=============================================")

	// var nonInitMap map[string]int
	// if _, ok = nonInitMap["somekey"]; !ok {
	// 	nonInitMap["somekey"] = 42
	// }
	// fmt.Println(nonInitMap["somekey"])

	initMap := map[string]int{}
	initMap["somekey"] = 12
	fmt.Println(initMap) // map[somekey:12]

	/************ delete element ***********/

	hp, ok = playerHPs["www"]
	fmt.Println(hp, ok)
	if ok {
		delete(playerHPs, "www")
	}
	fmt.Println(playerHPs)

	/************ iteration ***********/

	fmt.Println("=============================================")

	primes := map[int]bool{
		2: true,
		3: false,
		4: false,
		5: true,
	}

	fmt.Println(primes[2])  // true
	fmt.Println(primes[50]) // false

	for key, val := range primes {
		fmt.Println(key, "->", val)
		// ค่าจะไม่เรียงตาม key
	}
}
