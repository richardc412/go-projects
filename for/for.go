package main

import "fmt"

func main() {

	i := 1
	// basic loop, essentially a while condition
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// Standard for loop
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	// Create an iterator, loop
	for i := range 3 {
		fmt.Println(i)
	}

	// I think this should loop 3 times
	for range 3 {
		fmt.Println("This should show up 3 times")
	}

	// for can be used as a while true with a break line
	for {
		fmt.Println("Break loop")
		break
	}

	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}