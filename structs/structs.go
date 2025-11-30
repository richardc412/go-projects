package main

import (
	"fmt"
)

type person struct {
	name string
	age int
}

func newPerson(name string, age int) *person {
	p := person{name, age}
	return &p
}

func main() {
	var person1 *person = newPerson("Harry", 17)
	fmt.Println("Harry's name", person1.name)
	fmt.Println("Harry's age", person1.age)

	person1.name = "Bob"
	fmt.Println(person1)
}