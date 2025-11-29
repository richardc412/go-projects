package main

import "fmt"

func main() {
	/*
	Can put a statement before the conditional, and any variables declared in the statement are available
	throughout the entire branching statement (all ifs and elses)
	*/
	if num := 9; num < 0 {
		fmt.Println("num is smaller than 0")
	} else if num < 10 {
		fmt.Println("num is between 0 and 9")
	} else {
		fmt.Println("num is 10 or greater")
	}
}