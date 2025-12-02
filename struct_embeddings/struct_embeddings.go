package main

import (
	"fmt"
)

type base struct {
	num int
}

type embed struct {
	base
}

func (b base) describe() {
	fmt.Println(b.num)
}

func main() {
	e := embed{base: base{num: 3}}

	// The describe method of base gets promoted to embed
	e.describe()
	// This also works
	e.base.describe()
}