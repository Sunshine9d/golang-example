package main

import (
	"fmt"
	"github.com/Sunshine9d/golang-example/routine"
)

type NewAdvance struct {
	name string
}

func (a *NewAdvance) Run() {
	fmt.Println("Running")
}
func main() {
	// Create a new instance of the Advance struct
	ad := NewAdvance{}
	// Call the Run method
	ad.Run()
	routine.calculate()
}
