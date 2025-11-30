package main

import (
	"fmt"
)

type rect struct {
	width, height int
}

func createRect(w, h int) *rect {
	return &rect{w,h}
}

// takes a pointer to a rect, avoids copying and can mutate the struct
func (r *rect) area() int {
	return r.height * r.width
}

// takes a value of a struct, this value is copied from the struct passed to the function
func (r rect) perim() int {
	return 2*r.height + 2*r.width
}

// mutates a rect to double its height and width
func (r *rect) doubleSize() {
	r.height = 2*r.height
	r.width = 2*r.width
}

func main() {
	rect1 := createRect(5, 10)

	fmt.Println("Area:", rect1.area())
	fmt.Println("Perimeter", rect1.perim())

	rect1.doubleSize()

	fmt.Println("Area:", rect1.area())
	fmt.Println("Perimeter", rect1.perim())
}

