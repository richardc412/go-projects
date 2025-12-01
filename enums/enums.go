package main

// Implement states

import (
	"fmt"
)

type ToasterState int

const (
	Unplugged ToasterState = iota
	Ready
	Toasting
	Toasted
)

var toasterStateName = map[ToasterState]string {
	Unplugged: "Unplugged",
	Ready: "Ready to Use",
	Toasting: "Toasting Bread",
	Toasted: "Toasted",
}

func (ts ToasterState) String() string {
	return toasterStateName[ts]
}

func (ts *ToasterState) nextState() {
	switch *ts {
	case Unplugged:
		*ts = Ready
	case Ready:
		*ts = Toasting
	case Toasting:
		*ts = Toasted
	case Toasted:
		*ts = Ready
	}
}

func main() {
	curState := Unplugged
	fmt.Println(curState)
	curState.nextState();
	fmt.Println(curState)
	curState.nextState();
	fmt.Println(curState)
	curState.nextState();
	fmt.Println(curState)
	curState.nextState();
	fmt.Println(curState)
}