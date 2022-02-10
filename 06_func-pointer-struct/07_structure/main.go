package main

import (
	"fmt"
	"time"
)

type Person struct {
	name        string
	age         int
	dateOfBirth time.Time
}

func main() {
	var p1 Person
	p1.name = "Almond"
	p1.age = 18
	p1.dateOfBirth = time.Now()

	p2 := Person{
		name:        "Mango",
		age:         22,
		dateOfBirth: time.Now(),
	}

	p2Ptr := &p2
	p2Ptr.age = 30

	p2Copy := p2
	p2Copy.age = 40

	fmt.Println(p1.name, p2.name, p2.age, p2Copy.age)

	incAge(p1)
	fmt.Println(p1.age)

	incAgePtr(&p1)
	fmt.Println(p1.age)
}

func incAge(p Person) {
	p.age++
}

func incAgePtr(p *Person) {
	p.age++
}
