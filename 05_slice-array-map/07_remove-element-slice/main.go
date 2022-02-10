package main

import "fmt"

func main() {
	zombies := []string{"Paul", "Katya", "George", "Lucy"}
	fmt.Println("all zombies", zombies)
	toRemove := 2
	zombiesWithoutGeorge := append(zombies[:toRemove], zombies[toRemove+1:]...)
	// ให้ slice ตัวแรก zombies[:2] ก็คือ นับจาก index 0 ไปตำแหน่งที่ 2 จะได้ Paul, Katya
	// ให้ slice ตัวที่สอง zombies[2+1:] หรือ zombies[3:] ก็คือ นับจาก index 3 ไปยังตำแหน่งสุดท้าย จะได้แค่ Lucy

	fmt.Println("zombies without george", zombiesWithoutGeorge)

	toRemove = 1
	zombiesCopy := make([]string, len(zombies[:toRemove]))
	fmt.Println(zombiesCopy) // []
	copy(zombiesCopy, zombies[:toRemove])
	fmt.Println(zombiesCopy) // [Paul]
	zombiesWithoutKatya := append(zombiesCopy, zombies[toRemove+1:]...)
	fmt.Println("zombies without Katya", zombiesWithoutKatya) // [Paul Lucy]
}
