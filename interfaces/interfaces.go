package main

import (
	"fmt"
)

type thing interface {
	getName() string
}

type person struct {
	name string
}

type dog struct {
	name string
}

func (p *person) getName() string {
	return p.name
}

func (d dog) getName() string {
	return d.name
}

func getNameOfThing(t thing) string {
	return t.getName()
}

func main() {
	p := &person{"Adam"}
	person1 := person{"Bob"}
	dog := dog{"Rocket"}

	
	fmt.Println(getNameOfThing(&person1))

	fmt.Println(getNameOfThing(p))
	fmt.Println(getNameOfThing(dog))


}