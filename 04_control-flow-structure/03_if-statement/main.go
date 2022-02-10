package main

import (
	"fmt"
	"strings"
)

func main() {
	var prefix string
	fmt.Scan(&prefix)

	prefix = strings.TrimSpace(prefix)

	if prefix == "Mr." {
		fmt.Println("Hello Sir")
	} else if prefix == "Ms." {
		fmt.Println("Hello Madam")
	} else {
		fmt.Println("Hello there !!!F")
	}
}
