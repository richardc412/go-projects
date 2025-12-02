package main

import "fmt"

type Pair[T any] struct{
	first element[T]
	second element[T]
}

type element[T any] struct {
	id int
	val T
}

func (p Pair[T]) getSecondVal() T {
	return p.second.val
}

func main() {
	p := Pair[int]{first: element[int]{1, 10}, second: element[int]{2, 5}}
	secondVal := p.getSecondVal()
	fmt.Println(secondVal)
}