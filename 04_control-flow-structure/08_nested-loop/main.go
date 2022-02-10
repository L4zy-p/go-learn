package main

import "fmt"

func main() {
	for cnt := 0; cnt < 10; cnt++ {
		for cnt2 := 0; cnt2 < 10; cnt2++ {
			fmt.Println(cnt, cnt2)
		}
	}
}
