package main

import (
	"fmt"
)

func zeroVal(a int) {
	a = 0
}

func zeroPtr(a *int) {
	*a = 0
}

func main() {
	val := 1
	fmt.Println("val:", val)

	zeroVal(val)
	fmt.Println("zeroVal:", val)

	zeroPtr(&val)
	fmt.Println("zeroPtr", val)
}